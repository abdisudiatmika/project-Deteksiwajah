package controllers

import (
	"go-training-restful/config"
	"go-training-restful/database"
	"go-training-restful/models"
	"net/http"

	"github.com/labstack/echo"
)

var db = config.DB

//GetUsersController function mengambil data user
func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  true,
		"massage": "mendapatkan data",
		"data":    users,
	})
}

//CreateUserController function menambah user
func CreateUserController(user *models.User) (interface{}, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
