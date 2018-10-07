package models

// FaceCompResponse holds response from comparison
type FaceCompResponse struct {
	Confidence  float32 `json:"confidence"`
	IsIdentical bool    `json:"isIdentical"`
}
