package routes

import (
	c "FinalProjectQrCode/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)

	e.GET("/createqrcode", c.GetCreateQRCode)
	e.POST("/absensi/:time_stamp/:userid", c.AbsenController)

	e.POST("/addusers/", c.CreateUserController)
	e.POST("/addabsen/", c.CreateAbsenController)

	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id".UpdateUserController)

	eAuth := e.Group("")
	eAuth.Use(echoMid.BasicAuth(m.BasicAuth))
	eAuth.DELETE("/users/:id", c.DeleteUserController)
	eAuth.PUT("users/:id", c.UpdateUserController)
	return e
}
