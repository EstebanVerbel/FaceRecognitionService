package facecompserv

// API doc: https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395236

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Dev/FaceRecognitionService/models/constants"
)

// Detect faces on an image
func Detect() {

	// TODO:
	// * Take two images params to compare
	// * Implement call to cognitive service api
	// * return score

	const imageURL = "http://okmagazine.com/wp-content/uploads/2017/07/Angelina-Jolie-Bells-Palsy-Vanity-Fair-Interview-Long.jpg"
	//const imageURL = "https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/Angelina_Jolie_2_June_2014_%28cropped%29.jpg/220px-Angelina_Jolie_2_June_2014_%28cropped%29.jpg"

	const uri = constants.FaceDetectURIBase + constants.FaceDetectParams

	const imageURLEnc = "{\"url\":\"" + imageURL + "\"}"

	reader := strings.NewReader(imageURLEnc)

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
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", constants.SubscriptionKey1)

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the response body.
	// Note, data is a byte array
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the Json data
	var f interface{}
	json.Unmarshal(data, &f)

	// Format and display the Json result
	jsonFormatted, _ := json.MarshalIndent(f, "", "  ")

	fmt.Println(string(jsonFormatted))
}
