package repository

import (
	"gin-boiler/internal/database"
	"gin-boiler/internal/models"
)

func GetUser(id string) (*models.User, error) {
	User := &models.User{}
	err := database.DB.Where("id = ? AND account_status = ?", id, "active").First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func CreateUser(userModel models.User) error {
	err := database.DB.Create(&userModel).Error
	if err != nil {
		return err
	}
	return nil
}
