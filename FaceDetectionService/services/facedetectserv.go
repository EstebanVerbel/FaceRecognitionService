package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Dev/FaceRecognitionService/FaceDetectionService/models"
	"github.com/Dev/FaceRecognitionService/FaceDetectionService/models/constants"
)

// Detect detects face on image
func Detect(image []byte) models.FaceDetectResponse {

	const uri = constants.FaceDetectURIBase + constants.FaceDetectParams

	reader := bytes.NewBuffer(image)

	// Create the Http client
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	// Create the Post request, passing the image URL in the request body
	req, err := http.NewRequest("POST", uri, reader)
	if err != nil {
		panic(err)
	}

	req.Header.Add(constants.ContentTypeHeader, constants.BynaryContentType)
	req.Header.Add(constants.SubscriptionKeyHeader, constants.SubscriptionKey1)

	// Send the request and retrieve the response
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error on http request:", err)
	}

	defer resp.Body.Close()
	fmt.Println("Response Code: ", resp.StatusCode)

	// Read the response body.
	// Note, data is a byte array
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro while reading data from response: ", err)
	}

	var faceResponse models.FaceDetectResponse
	var detectionResponseArray []models.FaceDetectResponse

	unmarshalError := json.Unmarshal(data, &detectionResponseArray)
	if unmarshalError != nil {
		fmt.Println("Error unmarshaling json response", unmarshalError)
		// return
	}

	if len(detectionResponseArray) == 1 {

		fmt.Println("Face Detected Succesfully")

		faceResponse = detectionResponseArray[0]

		fmt.Println(faceResponse.FaceID)
		fmt.Println(faceResponse.FaceAttributes.Gender)

		return faceResponse
	} else if len(detectionResponseArray) > 1 {
		// error, more than one face on image
		return faceResponse
	} else {
		// No face on picture
		return faceResponse
	}
}
