package repository

import (
	"gin-boiler/internal/database"
	"gin-boiler/internal/models"
)

func GetListPost(offset int, limit int) ([]models.Post, int64, error) {
	var posts []models.Post
	var count int64

	if err := database.DB.Model(&models.Post{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Preload를 사용하여 User 관계를 미리 로드하고 페이징 적용
	// user는 active == true 인 데이터만
	// 즉 활성화 된 유저가 작성한 게시글만 불러옴
	err := database.DB.Preload("User", "active = ?", true).Offset(offset).Limit(limit).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}
