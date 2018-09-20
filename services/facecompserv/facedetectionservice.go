package facecompserv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// API doc: https://westus.dev.cognitive.microsoft.com/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395236

// Detect faces on an image
func Detect() {

	// TODO:
	// * Take two images params to compare
	// * Implement call to cognitive service api
	// * return score

	const subscriptionKey = "<Add Key Here>"

	// apiURL := URL + "/analyze?visualFeatures=" + query
	const uriBase = "https://eastus.api.cognitive.microsoft.com/face/v1.0/detect"

	const imageURL = "https://upload.wikimedia.org/wikipedia/commons/3/37/Dagestani_man_and_woman.jpg"

	const params = "?returnFaceAttributes=age,gender,headPose,smile,facialHair," +
		"glasses,emotion,hair,makeup,occlusion,accessories,blur,exposure,noise"

	const uri = uriBase + params
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
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

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
