package server

import (
	"basis/algorithm"
	"basis/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func Start() {
	http.HandleFunc("/submit-quiz", handleSubmittedQuiz)
	http.HandleFunc("/sign-in", signIn)
	http.HandleFunc("/sign-up", signUp)
	http.HandleFunc("/history", accessHistory)
	fmt.Println("The server is up")
	log.Fatalln(http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil))
}

func handleSubmittedQuiz(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	quizData, err := parseQuizData(req.URL)
	if err != nil {
		result, _ := json.Marshal("Incorrect input format")
		w.WriteHeader(400)
		w.Write(result)
		return
	}
	wardrobe, err := algorithm.CalculateWardrobe(quizData)
	var result []byte
	if err != nil {
		result, _ = json.Marshal("Internal error")
		w.WriteHeader(500)
	} else {
		result, _ = json.Marshal(wardrobe)
	}
	w.Write(result)
	tokenCookie, err := req.Cookie("access_token")
	if err != nil {
		return
	}
	username, err := validateTokenAndExtractUsername([]byte(tokenCookie.Value))
	if err != nil {
		return
	}
	db.SaveResult(wardrobe, username)
}

func accessHistory(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	tokenCookie, err := req.Cookie("access_token")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	username, err := validateTokenAndExtractUsername([]byte(tokenCookie.Value))
	if err != nil {
		w.WriteHeader(403)
		return
	}
	history, err := db.QueryResultsByDateByUsername(username)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var result []byte
	result, _ = json.Marshal(history)
	w.Write(result)
}

func parseQuizData(url *url.URL) (db.QuizData, error) {
	result := db.QuizData{}
	var err error
	result.Age, err = strconv.Atoi(url.Query().Get("age"))
	if err != nil {
		return db.QuizData{}, err
	}
	result.Budget, err = strconv.Atoi(url.Query().Get("budget"))
	if err != nil {
		return db.QuizData{}, err
	}
	result.Designation = url.Query().Get("designation")
	result.HairColor = url.Query().Get("hairColor")
	result.FavoriteColorScheme = url.Query().Get("favoriteColorScheme")
	result.PreferredFit = url.Query().Get("preferredFit")
	return result, nil
}
