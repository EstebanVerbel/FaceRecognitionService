package constants

// FaceDetectURIBase is the base uri for Azure's face detection service
const FaceDetectURIBase = "https://eastus.api.cognitive.microsoft.com/face/v1.0/detect?"

// FaceDetectParams holds the parameters wanted from the face detection service
const FaceDetectParams = "returnFaceAttributes=age,gender,headPose,smile"

// BynaryContentType contains string to define []byte content type on http requests
const BynaryContentType = "application/octet-stream"

// JSONContentType contains string to define Json content type on http requests
const JSONContentType = "application/json"

// ContentTypeHeader string Header
const ContentTypeHeader = "Content-Type"

// SubscriptionKeyHeader string header
const SubscriptionKeyHeader = "Ocp-Apim-Subscription-Key"
