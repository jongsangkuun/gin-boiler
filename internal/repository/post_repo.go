package repository

import (
	"gin-boiler/internal/database"
	"gin-boiler/internal/models"
)

func GetListPost(offset int, limit int) ([]models.Post, int64, error) {
	var posts []models.Post
	var count int64

	if err := database.RDB.Model(&models.Post{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := database.RDB.Offset(offset).Limit(limit).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}

func GetPostById(id string) (*models.Post, error) {
	post := &models.Post{}
	err := database.RDB.Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
