package models

import (
	"github.com/jinzhu/gorm"
)

//Usereye Struck simpen data user
type Usereye struct {
	gorm.Model

	Jarak    int    `json:"jarak" json:"jarak" query:"jarak"`
	Keyword  string `json:"keyword" json:"keyword" query:"keyword"`
	Personal string `json:"personal" json:"personal" query:"personal"`
}
