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
// @Success      200  {object}  utils.BaseResponse{data=dto.UserLoginResDto} "로그인 성공"
// @Failure      400  {object}  utils.BaseResponse "잘못된 요청"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /auth/login [post]
func UserLoginService(c *gin.Context) {
	var loginDto dto.UserLoginReqDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := repository.GetUserByUserId(loginDto.UserId)
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
	loginResponse := dto.UserLoginResDto{
		Token:    token,
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
	}

	response := utils.CreateBaseResponse(http.StatusOK, "로그인 성공", loginResponse)
	c.JSON(http.StatusOK, response)
}

// AdminLoginService godoc
// @Summary      관리자 로그인
// @Description  관리자 ID와 비밀번호로 로그인합니다
// @Tags         인증
// @Accept       json
// @Produce      json
// @Param        request body dto.AdminLoginReqDto true "로그인 정보"
// @Success      200  {object}  utils.BaseResponse{data=dto.AdminLoginResDto} "로그인 성공"
// @Failure      400  {object}  utils.BaseResponse "잘못된 요청"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /auth/login/admin [post]
func AdminLoginService(c *gin.Context) {
	var loginDto dto.AdminLoginReqDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	admin, err := repository.GetUserByUserId(loginDto.AdminId)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginDto.Password))
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := utils.GenerateJWT(admin.UserId, admin.Email, admin.Username)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	loginResponse := dto.AdminLoginResDto{
		Token:     token,
		AdminId:   admin.UserId,
		AdminName: admin.Username,
		Email:     admin.Email,
	}

	response := utils.CreateBaseResponse(http.StatusOK, "로그인 성공", loginResponse)
	c.JSON(http.StatusOK, response)
}
