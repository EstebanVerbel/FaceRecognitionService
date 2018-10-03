package facecompserv

// API doc: https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395236

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Dev/FaceRecognitionService/models/constants"
	"github.com/Dev/FaceRecognitionService/models/faceservice"
)

// Detect faces on an image
func Detect(image []byte) faceservice.FaceDetectResponse {

	const imageURL = "http://okmagazine.com/wp-content/uploads/2017/07/Angelina-Jolie-Bells-Palsy-Vanity-Fair-Interview-Long.jpg"
	//const imageURL = "https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Angelina_Jolie_2_June_2014_%28cropped%29.jpg/220px-Angelina_Jolie_2_June_2014_%28cropped%29.jpg"

	// C:\Angelina.jpg

	const uri = constants.FaceDetectURIBase + constants.FaceDetectParams

	// const imageURLEnc = "{\"url\":\"" + imageURL + "\"}"

	// reader := strings.NewReader(imageURLEnc)

	//reader := strings.NewReader(image)

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

	// Add headers
	//req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Ocp-Apim-Subscription-Key", constants.SubscriptionKey1)

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

	var faceResponse faceservice.FaceDetectResponse
	var detectionResponseArray []faceservice.FaceDetectResponse

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
