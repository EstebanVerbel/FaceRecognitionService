package models

// FaceDetectResponse defines response of face detection
type FaceDetectResponse struct {
	FaceAttributes FaceAttributes `json:"faceAttributes"`
	FaceID         string         `json:"faceId"`
	FaceRectangle  FaceRectangle  `json:"faceRectangle"`
}

// FaceRectangle defines rectangle that includes the face
type FaceRectangle struct {
	Height int32 `json:"height"`
	Left   int32 `json:"left"`
	Top    int32 `json:"top"`
	Width  int32 `json:"width"`
}

// HeadPose defines pose of the face
type HeadPose struct {
	Pitch float32 `json:"pitch"`
	Roll  float32 `json:"roll"`
	Yaw   float32 `json:"yaw"`
}

// FaceAttributes defines the attributes of the face
type FaceAttributes struct {
	Age      float32  `json:"age"`
	Gender   string   `json:"gender"`
	HeadPose HeadPose `json:"headPose"`
	Smile    float32  `json:"smile"`
}
