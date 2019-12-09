package database

import "FinalProjectQrCode/models"

func CreateQrCode (qrcode *models.QrCode) (interface{}, error) {
	if err := db.Create(&qrcode ).Error; err != nil {
		return nil, err
	}
	return qrcode , nil
}

func ReadQrCode (time_stamp string ) (interface{}, error) {
	qrCode := models.QrCode{}
	if err := db.Where("time_stamp=?", time_stamp).First(&qrCode).Error; err != nil {
		return nil, err
	}
	return qrCode , nil
}
