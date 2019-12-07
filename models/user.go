package models

import (
	"github.com/jinzhu/gorm"
)

//User Struck simpen data user
type User struct {
	gorm.Model

	Email    string `json:"email" json:"email" query:"email"`
	Name     string `json:"name" json:"name" query:"name"`
	Password string `json:"password" json:"password" query:"password"`
}
