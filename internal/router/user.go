package router

import "github.com/gin-gonic/gin"

// 사용자 관련 라우트 설정
func SetupUserRoutes(router *gin.Engine) {
	// API 그룹 생성
	api := router.Group("/api/v1")
	{
		// 사용자 관련 라우트
		users := api.Group("/users")
		{
			//users.POST("", createUser)       // POST /api/v1/users
			//users.GET("/:id", getUser)       // GET /api/v1/users/:id
			//users.GET("", listUsers)         // GET /api/v1/users
			//users.PUT("/:id", updateUser)    // PUT /api/v1/users/:id
			//users.DELETE("/:id", softDeleteUser) // DELETE /api/v1/users/:id
			//user.DELETE("", hardDeleteUsers)  // DELETE /api/v1/users
			users.GET("/ping", ping)
		}

	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
