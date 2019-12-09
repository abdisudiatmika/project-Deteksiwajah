package middleware

import (
	"project-Deteksiwajah/config"
	"project-Deteksiwajah/models"

	"github.com/labstack/echo"
)

//BasicAuthDB fungsi aut basi databse
func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.Face
	if err := db.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil

}
