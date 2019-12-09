package controllers

import (
	"FinalProjectQrCode/lib/database"
	"FinalProjectQrCode/models"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func AbsenController(c echo.Context) error {
	var time_stamp = (c.Param("time_stamp"))
	var userid = (c.Param("userid"))

	_, err := database.ReadQrCode(time_stamp)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "QR Code tidak valid",
		})
	}
	useridint, _ := strconv.Atoi(userid)
	absens := models.Absens{
		Model:           gorm.Model{},
		UserID:          useridint,
		QrCodeTimeStamp: time_stamp,
	}
	database.CreateAbsen(&absens)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success absensi",
		"users":   absens,
	})
}
