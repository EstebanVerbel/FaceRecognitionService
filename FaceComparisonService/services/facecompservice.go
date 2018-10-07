package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Dev/FaceRecognitionService/FaceComparisonService/models"
	"github.com/Dev/FaceRecognitionService/FaceComparisonService/models/constants"
)

// Compare compares two images of faces and returns a score of the likelihood that it's the same person
func Compare(faceID1 string, faceID2 string) models.FaceCompResponse {

	jsonRequest := "{ \"faceId1\": \"" + faceID1 + "\", \"faceId2\": \"" + faceID2 + "\" }"

	reader := strings.NewReader(jsonRequest)

	httpClient := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("POST", constants.FaceCompURIBase, reader)
	if err != nil {
		fmt.Printf("Error Creating http Request: %v\n", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", constants.SubscriptionKey1)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error Making http request: %v\n", err)
	}

	fmt.Println("Response Status Code: ", resp.StatusCode)
	fmt.Println("Request Message: ", http.StatusText(resp.StatusCode), ". ", resp.Status)

	defer resp.Body.Close()

	// Read the response body.
	// Note, data is a byte array
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var faceCompResponse models.FaceCompResponse
	json.Unmarshal(data, &faceCompResponse)

	return faceCompResponse
}
