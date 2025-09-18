package repository

import (
	"gin-boiler/internal/database"
	"gin-boiler/internal/models"
	"time"
)

func CreateAdmin(adminModel models.Admin) error {
	err := database.DB.Create(&adminModel).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAdmin(id string) (*models.Admin, error) {
	Admin := &models.Admin{}
	err := database.DB.Where("admin_id = ?", id).First(&Admin).Error
	if err != nil {
		return nil, err
	}

	return Admin, nil
}

func UpdateAdmin(adminModel models.Admin) error {
	err := database.DB.Save(&adminModel).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteAdmin(id string) error {
	err := database.DB.Model(&models.Admin{}).Where("admin_id = ?", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteHardAdmin(id string) error {
	err := database.DB.Where("admin_id = ?", id).Delete(&models.Admin{}).Error
	if err != nil {
		return err
	}

	return nil
}
