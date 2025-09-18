package main

import (
	"gin-boiler/internal/config"
	"gin-boiler/internal/database"
	"gin-boiler/internal/models"
	"gin-boiler/internal/utils"
	"log"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Panic(err)
	}
	_, err = database.Connect(env)
	if err != nil {
		log.Panic(err)
	}
	err = initAdmin()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Admin user created successfully.")
}

func initAdmin() error {
	var admin models.Admin
	hashPassword, err := utils.HashPassword("admin")
	if err != nil {
		return err
	}

	admin.AdminName = "admin"
	admin.Password = hashPassword
	admin.Email = "admin@admin.com"
	admin.AdminId = "admin"
	admin.Role = models.SuperAdmin

	err = database.DB.Create(&admin).Error
	if err != nil {
		return err
	}
	return nil
}
