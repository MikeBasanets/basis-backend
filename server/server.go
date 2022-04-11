package server

import (
	"basis/algorithm"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Start() {
	http.HandleFunc("/", mapRequest)
	fmt.Println("The server is up")
	log.Fatalln(http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil))
}

func mapRequest(w http.ResponseWriter, req *http.Request) {
	switch trimmedUrl := strings.TrimSuffix(req.URL.Path, "/"); trimmedUrl {
	case "/submit-quiz":
		handleSubmittedQuiz(w, req)
	default:
		respondError404(w, req)
	}
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
}

func parseQuizData(url *url.URL) (algorithm.QuizData, error) {
	result := algorithm.QuizData{}
	var err error
	result.Age, err = strconv.Atoi(url.Query().Get("age"))
	if err != nil {
		return algorithm.QuizData{}, err
	}
	result.Budget, err = strconv.Atoi(url.Query().Get("budget"))
	if err != nil {
		return algorithm.QuizData{}, err
	}
	result.Designation = url.Query().Get("designation")
	result.HairColor = url.Query().Get("hairColor")
	result.FavouriteColorScheme = url.Query().Get("favouriteColorScheme")
	result.PreferredFit = url.Query().Get("preferredFit")
	return result, nil
}

func respondError404(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Endpoint not found. Check to see if the submitted URL is correct.")
	w.WriteHeader(http.StatusNotFound)
}
