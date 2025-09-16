package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/models"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserService godoc
// @Summary      사용자 생성
// @Description  새로운 사용자 계정을 생성합니다
// @Tags         사용자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateUserReqDto true "사용자 생성 정보"
// @Success      200  {object}  utils.BaseResponse "사용자 생성 성공"
// @Failure      400  {object}  utils.BaseResponse "잘못된 요청"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /user [post]
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

// GetUserService godoc
// @Summary      사용자 조회
// @Description  사용자 ID로 사용자 정보를 조회합니다
// @Tags         사용자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "사용자 ID"
// @Success      200  {object}  utils.BaseResponse "사용자 조회 성공"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      404  {object}  utils.BaseResponse "사용자 없음"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /user/{id} [get]
func GetUserService(c *gin.Context) {
	id := c.Param("id")
	data, err := repository.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// UpdateUserService godoc
// @Summary      사용자 수정
// @Description  사용자 정보를 수정합니다
// @Tags         사용자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.UpdateUserReqDto true "사용자 수정 정보"
// @Success      200  {object}  utils.BaseResponse "사용자 수정 성공"
// @Failure      400  {object}  utils.BaseResponse "잘못된 요청"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /user [put]
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

// DeleteUserService godoc
// @Summary      사용자 삭제 (소프트 삭제)
// @Description  사용자를 소프트 삭제합니다
// @Tags         사용자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "사용자 ID"
// @Success      200  {object}  utils.BaseResponse "사용자 삭제 성공"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      404  {object}  utils.BaseResponse "사용자 없음"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /user/{id} [delete]
func DeleteUserService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

// ListUserService godoc
// @Summary      사용자 목록 조회
// @Description  모든 사용자 목록과 총 개수를 조회합니다
// @Tags         사용자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  utils.BaseListResponse "사용자 목록 조회 성공"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /user/list [get]
func ListUserService(c *gin.Context) {
	userList, count, err := repository.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseListResponse(http.StatusOK, "success", userList, count)
	c.JSON(http.StatusOK, response)

}

// DeleteHardUserService godoc
// @Summary      사용자 완전 삭제
// @Description  사용자를 데이터베이스에서 완전히 삭제합니다
// @Tags         사용자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "사용자 ID"
// @Success      200  {object}  utils.BaseResponse "사용자 완전 삭제 성공"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      404  {object}  utils.BaseResponse "사용자 없음"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /user/{id}/hard [delete]
func DeleteHardUserService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteUserHard(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}
