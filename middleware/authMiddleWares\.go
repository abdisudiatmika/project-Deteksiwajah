package middleware

import (
	"github.com/labstack/echo"
)

//BasicAuth mencoba basic autication
func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}
