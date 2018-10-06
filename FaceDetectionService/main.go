package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/Dev/FaceRecognitionService/FaceDetectionService/services"
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

	services.Detect(imageByteArray)

}
