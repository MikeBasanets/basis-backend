package server

import (
	"basis/db"
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey *rsa.PrivateKey

func InitJwtKey() {
	pemString := os.Getenv("JWT_KEY")
	pemString = strings.ReplaceAll(pemString, `\n`, "\n")
	block, _ := pem.Decode([]byte(pemString))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	jwtKey = key
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	err := json.NewDecoder(bytes.NewBuffer(bodyBytes)).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usernameValid, _ := regexp.MatchString("[a-zA-Z0-9]{3,}", creds.Username)
	passwordValid, _ := regexp.MatchString("[a-zA-Z0-9]{3,}", creds.Password)
	if !usernameValid || !passwordValid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, userNotFoundErr := db.QueryUserByUsername(creds.Username)
	if userNotFoundErr == nil {
		w.WriteHeader(http.StatusConflict) // username taken
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = db.SaveUser(db.User{Username: creds.Username, PasswordHash: string(hashedPassword), LastActive: time.Now()})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // username taken
		return
	}
	signIn(w, r)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedUser, userNotFoundErr := db.QueryUserByUsername(creds.Username)

	passwordMismatchErr := bcrypt.CompareHashAndPassword([]byte(expectedUser.PasswordHash), []byte(creds.Password))
	if userNotFoundErr != nil || passwordMismatchErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := jwt.NewBuilder().
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(5*time.Minute)).
		Claim("username", creds.Username).
		Build()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, jwtKey)) //os.Getenv("JWT_KEY")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   string(signedToken),
		Expires: time.Now().Add(100000 * time.Hour),
	})
}

func validateTokenAndExtractUsername(token []byte) (string, error) {
	verifiedToken, err := jwt.Parse(token, jwt.WithKey(jwa.RS256, jwtKey))
	if err != nil {
		return "", err
	}
	username := fmt.Sprintf("%v", verifiedToken.PrivateClaims()["username"])
	return username, nil
}
