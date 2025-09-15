package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserPingService(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateUserService(c *gin.Context) {
	var createDto dto.CreateUserDto
	if err := c.ShouldBindJSON(&createDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := repository.CreateUser(createDto)
	if err != nil {
		response := utils.CreateBaseResponse(500, "error", err.Error())
		c.JSON(500, response)
	}

	response := utils.CreateBaseResponse(200, "success", "")
	c.JSON(200, response)
}

func GetUserService(c *gin.Context) {
	id := c.Param("id")
	data, err := repository.GetUser(id)
	if err != nil {
		c.JSON(500, utils.CreateBaseResponse(500, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(200, "success", data)
	c.JSON(200, response)
}
func UpdateUserService(c *gin.Context)     {}
func DeleteUserService(c *gin.Context)     {}
func ListUserService(c *gin.Context)       {}
func DeleteHardUserService(c *gin.Context) {}
