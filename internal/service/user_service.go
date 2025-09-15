package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/models"
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
	// Request로 전달받은 데이터 중 비밀번호를 암호화는 로직이 Service or Repo(DB 통신 로직) 어디에 있어야 할까?
	hashPassword, err := utils.HashPassword(createDto.Password)
	if err != nil {
		response := utils.CreateBaseResponse(500, "error", err.Error())
		c.JSON(500, response)
	}

	// Dto -> Model
	userModel := models.User{
		Email:    createDto.Email,
		Password: hashPassword,
		UserId:   createDto.UserId,
		Username: createDto.Username,
	}

	err = repository.CreateUser(userModel)
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

func UpdateUserService(c *gin.Context) {
	var updateDto dto.UpdateUserDto
	var hashPassword string

	if err := c.ShouldBindJSON(&updateDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Dto -> Model
	userModel := models.User{
		Email:    updateDto.Email,
		Password: hashPassword,
		UserId:   updateDto.UserId,
		Username: updateDto.Username,
	}

	if updateDto.Password != "" {
		hashPassword, err := utils.HashPassword(updateDto.Password)
		if err != nil {
			response := utils.CreateBaseResponse(500, "error", err.Error())
			c.JSON(500, response)
		}
		userModel.Password = hashPassword
	}

	err := repository.UpdateUser(userModel)
	if err != nil {
		response := utils.CreateBaseResponse(500, "error", err.Error())
		c.JSON(500, response)
	}

	response := utils.CreateBaseResponse(200, "success", "")
	c.JSON(200, response)
}
func DeleteUserService(c *gin.Context)     {}
func ListUserService(c *gin.Context)       {}
func DeleteHardUserService(c *gin.Context) {}
