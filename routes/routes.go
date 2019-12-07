package routes

import (
	c "go-training-restful/controllers"

	"github.com/labstack/echo"
)

//New fungsi baru untuk routing
func New() *echo.Echo {
	e := echo.New()

	e.GET("users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)

	return e
}
