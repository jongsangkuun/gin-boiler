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
		SetupAuthRoutes(api)
	}
}
