package routes

import (
	c "project-Deteksiwajah/controllers"
	m "project-Deteksiwajah/middleware"

	"github.com/labstack/echo"
)

//New fungsi baru untuk routing
func New() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)
	e.POST("/userseye", c.CreatedbUsereyeController)

	e.POST("/findpersonal", c.FileUpload)

	e.GET("/userseye", c.GetUsersdbController)
	e.GET("/userseye/:id", c.GeteyedbbyidController)
	e.DELETE("/userseye/:id", c.DeleteeyedbController)
	e.PUT("/userseye/:id", c.UpdateeyedbController)

	// eUser := e.Group("/users/")
	// eUser.Use(echoMid.BasicAuth(m.BasicAuth))
	// eUser.PUT(":id", c.UpdateUserController)
	// eUser.DELETE(":id", c.DeleteUserController)
	//e.POST("findpersonal", c.GetFacePersonal)
	return e

}
