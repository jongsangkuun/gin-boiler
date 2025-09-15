package router

import (
	"gin-boiler/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(api *gin.RouterGroup) *gin.RouterGroup {
	userRoutes := api.Group("/user")
	{
		userRoutes.GET("/list", service.ListUserService)              // GET /api/v1/user
		userRoutes.GET("/:id", service.GetUserService)                // GET /api/v1/user/:id
		userRoutes.POST("", service.CreateUserService)                // POST /api/v1/user
		userRoutes.PUT("/", service.UpdateUserService)                // PUT /api/v1/user
		userRoutes.DELETE("/:id", service.DeleteUserService)          // DELETE /api/v1/user/:id
		userRoutes.DELETE("/:id/hard", service.DeleteHardUserService) // DELETE /api/v1/user
	}
	return userRoutes
}
