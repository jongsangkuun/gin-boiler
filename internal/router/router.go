package router

import (
	"github.com/gin-gonic/gin"
)

// 사용자 관련 라우트 설정
func SetupRoutes(router *gin.Engine) {
	// API 그룹 생성
	api := router.Group("/api/v1")
	{
		SetupUserRoutes(api)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
