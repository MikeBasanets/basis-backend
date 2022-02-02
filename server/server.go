package server

import (
	"basis/algorithm"
	"encoding/json"
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/", mapRequest)
	fmt.Println("The server is up")
    http.ListenAndServe(":8080", nil)
}

func mapRequest(w http.ResponseWriter, req *http.Request) {
	if (req.URL.Path == "/" || req.URL.Path == "") {
		handleStartPage(w, req)
		return
	}
	if (req.URL.Path == "/submit-data/" || req.URL.Path == "/submit-data") {
		handleSubmittedData(w, req)
		return
	}
	respondError404(w, req)
}

func handleStartPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, _ := json.Marshal("Hello")
	w.Write(data)
}

func handleSubmittedData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	userData := algorithm.UserData{}
	json.NewDecoder(req.Body).Decode(&userData)
	var wardrobe, err = algorithm.CalculateWardrobe(userData)
	var result []byte
	if err != nil {
		result, _ = json.Marshal(err)
	} else {
		result, _ = json.Marshal(wardrobe)
	}
	w.Write(result)
}

func respondError404(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Endpoint not found. Check to see if the submitted URL is correct.");
	w.WriteHeader(http.StatusNotFound)
}
