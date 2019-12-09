package models

//FaceWithLandmarks struck simpan data wajah
type FaceWithLandmarks struct {
	Successful bool `json:"Successful"`
	Faces      []struct {
		LeftEye []struct {
			X int `json:"X"`
			Y int `json:"Y"`
		} `json:"LeftEye"`
		RightEye []struct {
			X int `json:"X"`
			Y int `json:"Y"`
		} `json:"RightEye"`
	} `json:"Faces"`
}
