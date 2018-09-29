package faceservice

// FaceCompRequest defines faces comparison request body
type FaceCompRequest struct {
	FaceID1 string `json:"faceId1"`
	FaceID2 string `json:"faceid2"`
}

// FaceCompResponse holds response from comparison
type FaceCompResponse struct {
	Confidence  float32 `json:"confidence"`
	IsIdentical bool    `json:"isIdentical"`
}
