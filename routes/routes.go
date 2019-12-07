package routes

import (
	c "FinalProjectQrCode/controllers"
	m "FinalProjectQrCode/middleware"

	"github.com/labstack/echo"

	echoMid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	//e.DELETE("/users/:id", c.DeleteUserController)
	//e.PUT("/users/:id", c.UpdateUserController)

	eAuth := e.Group("")
	eAuth.Use(echoMid.BasicAuth(m.BasicAuth))
	eAuth.DELETE("/users/:id", c.DeleteUserController)
	eAuth.PUT("/users/:id", c.UpdateUserController)
	return e
}
