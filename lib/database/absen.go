package database

import "FinalProjectQrCode/models"

func CreateAbsen(absen *models.Absens) (interface{}, error) {
	if err := db.Create(&absen ).Error; err != nil {
		return nil, err
	}
	return absen , nil
}