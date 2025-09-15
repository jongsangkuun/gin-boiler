package router

import (
	"gin-boiler/internal/service"

	"github.com/gin-gonic/gin"
)

// 사용자 관련 라우트 설정
func SetupUserRoutes(router *gin.Engine) {
	// API 그룹 생성
	api := router.Group("/api/v1")
	{
		// Todo
		// 사용자 관련 라우트
		users := api.Group("/users")
		{
			users.GET("/ping", service.UserPingService)
			users.GET("/list", service.ListUserService)
			users.POST("", service.CreateUserService)           // POST /api/v1/users
			users.GET("/:id", service.GetUserService)           // GET /api/v1/users/:id
			users.PUT("/:id", service.UpdateUserService)        // PUT /api/v1/users/:id
			users.DELETE("/:id", service.DeleteUserService)     // DELETE /api/v1/users/:id
			users.DELETE("/:id", service.DeleteHardUserService) // DELETE /api/v1/users
		}
		//posts := api.Group("/posts")
		//{
		//	posts.GET("/ping", ping)
		//}
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
