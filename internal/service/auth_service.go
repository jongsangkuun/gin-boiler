package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(c *gin.Context) {
	var loginDto dto.LoginReqDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := repository.GetUserById(loginDto.UserId)
	if err != nil {
		response := utils.CreateBaseResponse(500, "error", err.Error())
		c.JSON(500, response)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		response := utils.CreateBaseResponse(500, "error", err.Error())
		c.JSON(500, response)
		return
	}

	token, err := utils.GenerateJWT(user.UserId, user.Email, user.Username)
	if err != nil {
		response := utils.CreateBaseResponse(500, "error", err.Error())
		c.JSON(500, response)
		return
	}
	loginResponse := dto.LoginResDto{
		Token:    token,
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
	}

	response := utils.CreateBaseResponse(http.StatusOK, "로그인 성공", loginResponse)
	c.JSON(http.StatusOK, response)
}
