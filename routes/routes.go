package routes

import (
	c "project-Deteksiwajah/controllers"

	"github.com/labstack/echo"
)

//New fungsi baru untuk routing
func New() *echo.Echo {
	e := echo.New()

	//m.LogMiddleware(e)
	e.POST("/userseye", c.CreatedbUsereyeController)

	e.POST("/findpersonal", c.FileUpload)

	e.GET("/userseye", c.GetUsersdbController)
	e.GET("/userseye/:id", c.GeteyedbbyidController)
	e.DELETE("/userseye/:id", c.DeleteeyedbController)
	e.PUT("/userseye/:id", c.UpdateeyedbController)

	return e

}
