package main

import (
	"github.com/Dev/FaceRecognitionService/services/facecompserv"
)

func main() {

	// TODO:
	// Run face detection on first image
	// Run face detection on second image
	// Run face comparison on both images
	// return score

	// var f faceservice.FaceCompResponse

	// const faceID1 string = "8f1f117b-809a-45d3-8b1d-fcc83f88b7d9"
	// const faceID2 string = "da857626-1ed4-408c-9fc9-378e86147907"

	facecompserv.Detect()

	// f = facecompserv.Compare(faceID1, faceID2)

	// fmt.Println("Printing result from Main")
	// fmt.Println("Confidence: ", f.Confidence)
	// fmt.Println("Is Identical: ", f.IsIdentical)
}
