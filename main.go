package main

import (
	"fmt"

	"github.com/Dev/FaceRecognitionService/models/faceservice"
	"github.com/Dev/FaceRecognitionService/services/facecompserv"
)

func main() {

	// run face detection
	//facecompserv.Detect()

	var f faceservice.FaceCompResponse
	f = facecompserv.Compare()

	fmt.Println("Printing result from Main")
	fmt.Println("Confidence: ", f.Confidence)
	fmt.Println("Is Identical: ", f.IsIdentical)
}
