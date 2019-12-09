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

//GetUsersdbController function mengambil data user
func GetUsersdbController(c echo.Context) error {
	users, err := database.Geteyedb()

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

//CreatedbUsereyeController function menambah user
func CreatedbUsereyeController(c echo.Context) error {
	user := models.Usereye{}

	c.Bind(&user)
	result, err := database.CreatedbUsereye(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result)

}

//GeteyedbbyidController function mengambil data user
func GeteyedbbyidController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := database.Geteyedbbyid(id)

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

// UpdateeyedbController Update by id
func UpdateeyedbController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.Usereye{}
	newUser := user
	c.Bind(&newUser)

	result, err := database.Updateeyedb(id, &newUser)

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

//DeleteeyedbController haous data
func DeleteeyedbController(c echo.Context) error {
	user := models.Usereye{}
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))
	result, err := database.Deleteeyedb(id)

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

	resultkey, resultpersonal, err := database.FindPersonal(jarak)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//fmt.Println(personal)
	var personal models.Personality
	personal.UserPersonality = resultpersonal
	fmt.Println(personal)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"status":  true,
		"massage": "mendapatkan data",
		"data":    GetQutes(resultkey, resultpersonal),
	})

}
