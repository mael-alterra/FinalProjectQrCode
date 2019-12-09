package controllers

import (
	"FinalProjectQrCode/lib/database"
	"github.com/labstack/echo"
	"net/http"
)

func AbsenController (c echo.Context) error {
	var time_stamp = (c.Param("time_stamp"))
	users, err := database.ReadQrCode(time_stamp)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "QR Code tidak valid",
		})
	}


	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success absensi",
		"users":   users,
	})
}
