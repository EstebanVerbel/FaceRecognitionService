package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/Dev/FaceRecognitionService/services/facecompserv"

	_ "image/jpeg"
)

func main() {

	imageFile, err := os.Open("Angelina.jpg")

	if err != nil {
		// error
		fmt.Println("Error reading image: ", err)
	}

	defer imageFile.Close()

	imageFile.Seek(0, 0)

	imageData, imageType, err := image.Decode(imageFile)
	if err != nil {
		// Handle error
		fmt.Println("Error Decoding Image: ", err)
	}

	fmt.Println(imageType)

	buf := new(bytes.Buffer)

	err = jpeg.Encode(buf, imageData, nil)
	if err != nil {
		fmt.Println("Error Encodimg image into byte: ", err)
	}

	imageByteArray := buf.Bytes()

	facecompserv.Detect(imageByteArray)

	fmt.Println("done")

	// TODO:
	// Run face detection on first image
	// Run face detection on second image
	// Run face comparison on both images
	// return score

	// var f faceservice.FaceCompResponse

	// const faceID1 string = "8f1f117b-809a-45d3-8b1d-fcc83f88b7d9"
	// const faceID2 string = "da857626-1ed4-408c-9fc9-378e86147907"

	// facecompserv.Detect()

	// f = facecompserv.Compare(faceID1, faceID2)

	// fmt.Println("Printing result from Main")
	// fmt.Println("Confidence: ", f.Confidence)
	// fmt.Println("Is Identical: ", f.IsIdentical)
}
