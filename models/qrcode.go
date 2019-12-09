package models

import "github.com/jinzhu/gorm"

type QrCode struct {
	gorm.Model
	TimeStamp   string    `json:"userid" 	form:"userid"`
}