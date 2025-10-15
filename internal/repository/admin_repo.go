package repository

import (
	"gin-boiler/internal/database"
	"gin-boiler/internal/models"
	"time"
)

func CreateAdmin(adminModel models.Admin) error {
	err := database.RDB.Create(&adminModel).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAdmin(id string) (*models.Admin, error) {
	Admin := &models.Admin{}
	err := database.RDB.Where("admin_id = ?", id).First(&Admin).Error
	if err != nil {
		return nil, err
	}

	return Admin, nil
}

func UpdateAdmin(adminModel models.Admin) error {
	err := database.RDB.Save(&adminModel).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteAdmin(id string) error {
	err := database.RDB.Model(&models.Admin{}).Where("admin_id = ?", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteHardAdmin(id string) error {
	err := database.RDB.Where("admin_id = ?", id).Delete(&models.Admin{}).Error
	if err != nil {
		return err
	}

	return nil
}
