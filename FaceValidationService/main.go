package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dev/FaceRecognitionService/FaceValidationService/models"
	"github.com/gorilla/mux"
)

func main() {

	// TODO: needs to create rest api that takes []byte

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", serviceIsRunning)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func serviceIsRunning(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ServiceRunningResponse{Status: "OK", Code: 200})
}
