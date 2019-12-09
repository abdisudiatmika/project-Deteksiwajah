package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"project-Deteksiwajah/database"
	"project-Deteksiwajah/models"
	"strconv"

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
	user := models.Usereye{}

	c.Bind(&user)
	result, err := database.CreateUsereye(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result)

}

//GetUserController function mengambil data user
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := database.GetUserByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  true,
		"massage": "mendapatkan data",
		"data":    result,
	})
}

// UpdateUserController Update by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.Usereye{}
	newUser := user
	c.Bind(&newUser)

	result, err := database.UpdateUserByID(id, &newUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  true,
		"massage": "Berhasil Update data",
		"data":    result,
	})

}

//DeleteUserController haous data
func DeleteUserController(c echo.Context) error {
	user := models.Usereye{}
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))
	result, err := database.DeleteUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  true,
		"massage": "Berhasil Hapus data",
		"data":    result,
	})

}

//FileUpload data
func FileUpload(c echo.Context) error {

	//Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./images/1.jpeg")
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return GetFacePersonal(c)
}

//GetFacePersonal mencari personal
func GetFacePersonal(c echo.Context) error {

	jarak := Getdataface()

	result, err := database.FindPersonal(jarak)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(result)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  true,
		"massage": "mendapatkan data",
		"data":    GetQutes(result),
	})

}
