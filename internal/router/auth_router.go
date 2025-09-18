package router

import (
	"gin-boiler/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(api *gin.RouterGroup) (a *gin.RouterGroup) {
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/login", service.UserLoginService)
	}
	return authRoutes
}
