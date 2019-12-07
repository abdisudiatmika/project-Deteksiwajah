package controllers

import (
	"go-training-restful/database"
	"go-training-restful/models"
	"net/http"

	"github.com/labstack/echo"
)

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
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	result, err := database.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result)

}
