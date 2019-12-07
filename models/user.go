package models

import "github.com/jinzhu/gorm"

//Users ...
type Users struct {
	gorm.Model
	Fullname string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
