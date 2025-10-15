package router

import (
	"gin-boiler/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(api *gin.RouterGroup) *gin.RouterGroup {
	postRoutes := api.Group("/post")
	{
		postRoutes.GET("/list", service.ListPostService)
		postRoutes.GET("/:id", service.GetPostService)
		postRoutes.POST("", service.CreatePostService)
		postRoutes.PUT("/:id", service.UpdatePostService)
		postRoutes.DELETE("/:id", service.DeletePostService)
		postRoutes.DELETE("/:id/hard", service.DeleteHardPostService)
	}
	return postRoutes
}
