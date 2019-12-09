package models

//Responface struck simpan data wajah
type Responface struct {
	Successful bool
	Faces      []Face

	// LeftEye   []int    `json:"lefteye"`
	// RightEyes []string `json:"RightEye"`
}
