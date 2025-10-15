package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPostService(c *gin.Context) {
	var pageDto dto.BasePageReqDto

	if err := c.ShouldBindJSON(&pageDto); err != nil {
		response := dto.CreateBaseListResponse(http.StatusBadRequest, "fail", err.Error(), 0)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	postData, count, err := repository.GetListPost(pageDto.Limit, pageDto.Offset)
	if err != nil {
		response := dto.CreateBaseListResponse(http.StatusInternalServerError, "error", err.Error(), 0)
		c.JSON(http.StatusInternalServerError, response)
	}

	response := dto.CreateBaseListResponse(http.StatusOK, "success", postData, count)
	c.JSON(http.StatusOK, response)
}

func GetPostService(c *gin.Context) {

}
func CreatePostService(c *gin.Context)     {}
func UpdatePostService(c *gin.Context)     {}
func DeletePostService(c *gin.Context)     {}
func DeleteHardPostService(c *gin.Context) {}
