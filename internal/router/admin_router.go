package router

import (
	"gin-boiler/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(api *gin.RouterGroup) *gin.RouterGroup {
	adminRouter := api.Group("/admin")
	{
		adminRouter.GET("/:id", service.GetAdminService)
		adminRouter.POST("", service.CreateAdminService)
		adminRouter.PUT("/", service.UpdateAdminService)
		adminRouter.DELETE("/:id", service.DeleteAdminService)
		adminRouter.DELETE("/:id/hard", service.DeleteHardAdminService)
	}
	return adminRouter
}
