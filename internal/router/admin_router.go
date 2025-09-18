package router

import (
	"gin-boiler/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(api *gin.RouterGroup) *gin.RouterGroup {
	adminRouter := api.Group("/admin")
	{
		adminRouter.GET("/:id", service.GetAdminService)                // GET /api/v1/admin/:id
		adminRouter.POST("", service.CreateAdminService)                // POST /api/v1/admin
		adminRouter.PUT("/", service.UpdateAdminService)                // PUT /api/v1/admin
		adminRouter.DELETE("/:id", service.DeleteAdminService)          // DELETE /api/v1/admin/:id
		adminRouter.DELETE("/:id/hard", service.DeleteHardAdminService) // DELETE /api/v1/admin
	}
	return adminRouter
}
