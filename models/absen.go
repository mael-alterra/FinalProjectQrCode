package models

import "github.com/jinzhu/gorm"

//Absen ...
type Absens struct {
	gorm.Model
	UserID   			int    `json:"userid" 	form:"userid"`
	QrCodeTimeStamp  	string `json:"qrcodetimestamp" form:"qrcodetimestamp"`
	//Tanggal  string `json:"tanggal" form:"tanggal`
	//Name     string `json:"name"	form:"name"`
	//Status   bool 	`json:"status"	form:"status"`
}