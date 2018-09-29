package constants

// FaceDetectURIBase is the base uri for Azure's face detection service
const FaceDetectURIBase = "https://eastus.api.cognitive.microsoft.com/face/v1.0/detect?"

// FaceDetectParams holds the parameters wanted from the face detection service
const FaceDetectParams = "returnFaceAttributes=age,gender,headPose,smile"

// FaceCompURIBase is the base uri for Azure's face verification service
const FaceCompURIBase = "https://eastus.api.cognitive.microsoft.com/face/v1.0/verify?"
