package repository

import (
	"gin-boiler/internal/database"
	"gin-boiler/internal/dto"
	"gin-boiler/internal/models"
)

func GetUser(id string) (*models.User, error) {
	db := database.DB
	User := &models.User{}
	err := db.Where("id = ? AND account_status = ?", id, "active").First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func CreateUser(user dto.CreateUserDto) error {
	db := database.DB
	userModel := models.User{
		Email:    user.Email,
		Password: user.Password,
		UserId:   user.UserId,
		Username: user.Username,
	}

	err := db.Create(&userModel).Error
	if err != nil {
		return err
	}
	return nil
}
