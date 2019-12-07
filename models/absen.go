package models

import "github.com/jinzhu/gorm"

//Absen ...
type Absen struct {
	gorm.Model
	Tanggal  string `json:"tanggal" form:"tanggal"`
	UserId   int    `json:"userid" form:"userid"`
	Password string `json:"password" form:"password"`
}
