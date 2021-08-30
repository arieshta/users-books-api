package database

import (
	"users-api/config"
	"users-api/models"
)

func CreateUser(user *models.Users) error {
	if err := config.DB.Table("users").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() (interface{}, error) {
	var users []models.Users

	if err := config.DB.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var user models.Users

	if err := config.DB.Table("users").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserById(id int, user *models.Users) error {
	var users models.Users
	if err := config.DB.Table("users").First(&users, id).Error; err != nil {
		return err
	}
	err := config.DB.Table("users").Where("id = ?", id).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int) error {
	var user models.Users
	if err := config.DB.Table("users").Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
