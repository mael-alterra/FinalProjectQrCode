package database

import (
	"FinalProjectQrCode/config"
	"FinalProjectQrCode/models"
)

var (
	db = config.DB
)

//GetUsers ...
func GetUsers() (interface{}, error) {
	var users []models.Users

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

//GetUsers ...
func GetUser(id int) (interface{}, error) {
	var user []models.Users

	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// create new user
func CreateUser(user *models.Users) (interface{}, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// delete user by id
func DeleteUser(id int) (interface{}, error) {
	var user, deletedUser models.Users
	db.First(&user, id)
	db.First(&deletedUser, id)

	if err := db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return deletedUser, nil
}

// update user by id
func UpdateUser(id int, newUser *models.Users) (interface{}, error) {
	user := models.Users{}
	if err := db.Where("ID = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	if err := db.Save(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}
