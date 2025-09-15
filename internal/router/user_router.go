package router

import (
	"gin-boiler/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(api *gin.RouterGroup) *gin.RouterGroup {
	userRoutes := api.Group("/user")
	{
		userRoutes.GET("/ping", service.UserPingService)              // GET /api/v1/user_routes/ping
		userRoutes.GET("/list", service.ListUserService)              // GET /api/v1/user_routes
		userRoutes.GET("/:id", service.GetUserService)                // GET /api/v1/user_routes/:id
		userRoutes.POST("", service.CreateUserService)                // POST /api/v1/user_routes
		userRoutes.PUT("/:id", service.UpdateUserService)             // PUT /api/v1/user_routes/:id
		userRoutes.DELETE("/:id", service.DeleteUserService)          // DELETE /api/v1/user_routes/:id
		userRoutes.DELETE("/:id/hard", service.DeleteHardUserService) // DELETE /api/v1/user_routes
	}
	return userRoutes
}
