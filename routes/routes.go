package routes

import (
	c "go-training-restful/controllers"

	"github.com/labstack/echo"
)

//New fungsi baru untuk routing
func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", c.GetUsersController)
	return e
}
