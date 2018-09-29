package main

import (
	"fmt"

	"github.com/Dev/FaceRecognitionService/models/faceservice"
	"github.com/Dev/FaceRecognitionService/services/facecompserv"
)

func main() {

	// TODO:
	// Run face detection on first image
	// Run face detection on second image
	// Run face comparison on both images
	// return score

	var f faceservice.FaceCompResponse
	f = facecompserv.Compare()

	fmt.Println("Printing result from Main")
	fmt.Println("Confidence: ", f.Confidence)
	fmt.Println("Is Identical: ", f.IsIdentical)
}
