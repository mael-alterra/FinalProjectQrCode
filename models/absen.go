package models

import "github.com/jinzhu/gorm"

//Absen ...
type absen struct {
	gorm.Model
	UserId   int    `json:"userid" form:"userid"`
	Tanggal  string `json:"tanggal" form:"tanggal
}

type user struct{
	gorm.Model
	absen absen `gorm:"Foreignkey:UserId"`
}
