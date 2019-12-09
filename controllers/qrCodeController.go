package controllers

import (
	"FinalProjectQrCode/lib/database"
	"FinalProjectQrCode/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	rqrcode "github.com/skip2/go-qrcode"
	wqrcode "github.com/tuotoo/qrcode"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//GetReadQRCode ...
func GetReadQRCode() {
	fi, err := os.Open("qr.png")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	qrmatrix, err := wqrcode.Decode(fi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(qrmatrix.Content)
}

//GetCreateQRCode creates QRCode
func GetCreateQRCode(c echo.Context) error {
	var timeStamp string = strconv.Itoa(time.Now().Nanosecond())
	fmt.Println(timeStamp)
	qrCode := models.QrCode{
		Model:     gorm.Model{},
		TimeStamp: timeStamp,
	}

	if _,err := database.CreateQrCode(&qrCode); err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := rqrcode.WriteFile("localhost:8000/absensi/"+timeStamp, rqrcode.Medium, 256, "qrCode.png")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get create QR Code",
		"url QR Code " : "C:/Users/Roby/go/src/FinalProjectQrCode/qrCode.png",
	})
}

func ReadAbsenQrCode ( c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param(""))
	users, err := database.GetUser(id)
	if err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	"message": "success get all users",
	"users":   users,
	})

}