package server

import (
	"basis/algorithm"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Start() {
	http.HandleFunc("/", mapRequest)
	fmt.Println("The server is up")
	go http.ListenAndServe(":80", nil)
	log.Fatal(http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil))
}

func mapRequest(w http.ResponseWriter, req *http.Request) {
	switch cleanedUrl := strings.TrimSuffix(req.URL.Path, "/"); cleanedUrl {
	case "":
		handleStartPage(w, req)
	case "/submit-quiz":
		handleSubmittedQuiz(w, req)
	}
	respondError404(w, req)
}

func handleStartPage(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/main.html")
}

func handleSubmittedQuiz(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	quizBody := algorithm.QuizData{}
	json.NewDecoder(req.Body).Decode(&quizBody)
	if (quizBody == algorithm.QuizData{}) {
		fmt.Println(req.URL.Query().Get("birthdayYear"))
		birthdayYear, err := strconv.Atoi(req.URL.Query().Get("birthdayYear"))
		if err != nil {
			quizBody = algorithm.QuizData{Purpose: req.URL.Query().Get("purpose"), BirthdayYear: birthdayYear}
		}
	}
	var wardrobe, err = algorithm.CalculateWardrobe(quizBody)
	var result []byte
	if err != nil {
		result, _ = json.Marshal("Bad request")
		w.WriteHeader(400)
	} else {
		result, _ = json.Marshal(wardrobe)
	}
	w.Write(result)
}

func respondError404(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w, "Endpoint not found. Check to see if the submitted URL is correct.")
	w.WriteHeader(http.StatusNotFound)
}
