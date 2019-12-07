package middleware

import (
	"go-training-restful/config"
	"go-training-restful/models"

	"github.com/labstack/echo"
)

//BasicAuthDB fungsi aut basi databse
func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	if err := db.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil

}
