package repository

import (
	"gin-boiler/internal/database"
	"gin-boiler/internal/models"
	"time"
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

func UpdateUser(userModel models.User) error {
	err := database.DB.Save(&userModel).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	err := database.DB.Model(&models.User{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserHard(id string) error {
	err := database.DB.Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserList() ([]models.User, int64, error) {
	var users []models.User
	var count int64

	err := database.DB.Find(&users).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
