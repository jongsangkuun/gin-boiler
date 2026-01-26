package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserLoginService godoc
// @Summary      사용자 로그인
// @Description  사용자 ID와 비밀번호로 로그인합니다
// @Tags         인증
// @Accept       json
// @Produce      json
// @Param        request body dto.UserLoginReqDto true "로그인 정보"
// @Success      200  {object}  dto.BaseResponseDto{data=dto.UserLoginResDto} "로그인 성공"
// @Failure      400  {object}  dto.BaseResponseDto "잘못된 요청"
// @Failure      401  {object}  dto.BaseResponseDto "인증 실패"
// @Failure      500  {object}  dto.BaseResponseDto "서버 오류"
// @Router       /auth/login [post]
func UserLoginService(c *gin.Context) {
	var loginDto dto.UserLoginReqDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		response := dto.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := repository.GetUserByUserId(loginDto.UserId)
	if err != nil {
		response := dto.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		response := dto.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := utils.GenerateJWT(user.LoginId, user.Email, user.NickName)
	if err != nil {
		response := dto.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	loginResponse := dto.UserLoginResDto{
		Token:    token,
		UserId:   user.LoginId,
		Username: user.NickName,
		Email:    user.Email,
	}

	response := dto.CreateBaseResponse(http.StatusOK, "로그인 성공", loginResponse)
	c.JSON(http.StatusOK, response)
}

// AdminLoginService godoc
// @Summary      관리자 로그인
// @Description  관리자 ID와 비밀번호로 로그인합니다
// @Tags         인증
// @Accept       json
// @Produce      json
// @Param        request body dto.AdminLoginReqDto true "로그인 정보"
// @Success      200  {object}  dto.BaseResponseDto{data=dto.AdminLoginResDto} "로그인 성공"
// @Failure      400  {object}  dto.BaseResponseDto "잘못된 요청"
// @Failure      401  {object}  dto.BaseResponseDto "인증 실패"
// @Failure      500  {object}  dto.BaseResponseDto "서버 오류"
// @Router       /auth/login/admin [post]
func AdminLoginService(c *gin.Context) {
	var loginDto dto.AdminLoginReqDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		response := dto.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	admin, err := repository.GetAdmin(loginDto.AdminId)
	if err != nil {
		response := dto.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginDto.Password))
	if err != nil {
		response := dto.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := utils.GenerateJWT(admin.LoginId, admin.Email, admin.NickName)
	if err != nil {
		response := dto.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	loginResponse := dto.AdminLoginResDto{
		Token:     token,
		AdminId:   admin.LoginId,
		AdminName: admin.NickName,
		Email:     admin.Email,
	}

	response := dto.CreateBaseResponse(http.StatusOK, "로그인 성공", loginResponse)
	c.JSON(http.StatusOK, response)
}
