package models

import (
	"github.com/jinzhu/gorm"
)

//Usereye Struck simpen data user
type Usereye struct {
	gorm.Model

	Jarak    int    `json:"jarak" json:"jarak" query:"jarak"`
	Personal string `json:"personal" json:"personal" query:"personal"`
	//Password string `json:"password" json:"password" query:"password"`
}
