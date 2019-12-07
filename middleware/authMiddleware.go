package middleware

import (
	"fmt"
	"github.com/labstack/echo"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	fmt.Println(username, password)
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}
