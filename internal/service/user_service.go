package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/models"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserService(c *gin.Context) {
	var createDto dto.CreateUserReqDto
	if err := c.ShouldBindJSON(&createDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Request로 전달받은 데이터 중 비밀번호를 암호화는 로직이 Service or Repo(DB 통신 로직) 어디에 있어야 할까?
	hashPassword, err := utils.HashPassword(createDto.Password)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
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
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

func GetUserService(c *gin.Context) {
	id := c.Param("id")
	data, err := repository.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func UpdateUserService(c *gin.Context) {
	var updateDto dto.UpdateUserReqDto
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
			response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
			c.JSON(http.StatusInternalServerError, response)
		}
		userModel.Password = hashPassword
	}

	err := repository.UpdateUser(userModel)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

func DeleteUserService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

func ListUserService(c *gin.Context) {
	userList, count, err := repository.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseListResponse(http.StatusOK, "success", userList, count)
	c.JSON(http.StatusOK, response)

}

func DeleteHardUserService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteUserHard(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}
