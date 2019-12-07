package middleware

import (
	"FinalProjectQrCode/config"
	"FinalProjectQrCode/models"

	"github.com/labstack/echo"
)

var db = config.DB

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var user models.Users
	if err := db.Where("email = ? and password = ?", email, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
