package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginService godoc
// @Summary      사용자 로그인
// @Description  사용자 ID와 비밀번호로 로그인합니다
// @Tags         인증
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginReqDto true "로그인 정보"
// @Success      200  {object}  utils.BaseResponse{data=dto.LoginResDto} "로그인 성공"
// @Failure      400  {object}  utils.BaseResponse "잘못된 요청"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /auth/login [post]
func LoginService(c *gin.Context) {
	var loginDto dto.LoginReqDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := repository.GetUserById(loginDto.UserId)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := utils.GenerateJWT(user.UserId, user.Email, user.Username)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
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
