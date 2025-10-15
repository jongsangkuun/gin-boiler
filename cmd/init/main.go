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
		log.Println(err)
	}
	err = initUser()
	if err != nil {
		log.Println(err)
	}
	log.Println("Admin & User created successfully.")
}

func initAdmin() error {
	var admin models.Admin
	hashPassword, err := utils.HashPassword("admin")
	if err != nil {
		return err
	}

	admin.NickName = "admin"
	admin.Password = hashPassword
	admin.Email = "admin@admin.com"
	admin.LoginId = "admin"
	admin.Role = models.SuperAdmin

	err = database.DB.Create(&admin).Error
	if err != nil {
		return err
	}
	return nil
}

func initUser() error {
	var user models.User
	hashPassword, err := utils.HashPassword("user")
	if err != nil {
		return err
	}

	user.NickName = "user"
	user.Password = hashPassword
	user.Email = "user@user.com"
	user.LoginId = "user"
	user.AccountStatus = models.AccountStatusActive
	err = database.DB.Create(&user).Error

	if err != nil {
		return err
	}
	return nil
}
