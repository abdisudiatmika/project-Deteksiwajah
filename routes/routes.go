package routes

import (
	c "go-training-restful/controllers"
	m "go-training-restful/middleware"

	"github.com/labstack/echo"
)

//New fungsi baru untuk routing
func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)
	e.GET("users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)

	// eUser := e.Group("/users/")
	// eUser.Use(echoMid.BasicAuth(m.BasicAuth))
	// eUser.PUT(":id", c.UpdateUserController)
	// eUser.DELETE(":id", c.DeleteUserController)

	return e

}
