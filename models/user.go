package models

import (
	"github.com/jinzhu/gorm"
)

//Personal Struck simpen data user
type Personal struct {
	gorm.Model

	Jarak    string `json:"jarak" json:"jarak" query:"jarak"`
	Personal string `json:"personal" json:"personal" query:"personal"`
	//Password string `json:"password" json:"password" query:"password"`
}
