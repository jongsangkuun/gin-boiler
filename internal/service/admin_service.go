package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/models"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAdminService godoc
// @Summary      관리자 생성
// @Description  새로운 관리자 계정을 생성합니다
// @Tags         관리자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateAdminReqDto true "관리자 생성 정보"
// @Success      200  {object}  utils.BaseResponse "관리자 생성 성공"
// @Failure      400  {object}  utils.BaseResponse "잘못된 요청"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /admin [post]
func CreateAdminService(c *gin.Context) {
	var createDto dto.CreateAdminReqDto
	if err := c.ShouldBindJSON(&createDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
	}

	hashPassword, err := utils.HashPassword(createDto.Password)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	}

	adminModel := models.Admin{
		Email:     createDto.Email,
		Password:  hashPassword,
		AdminId:   createDto.AdminId,
		AdminName: createDto.AdminName,
	}

	err = repository.CreateAdmin(adminModel)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

// GetAdminService godoc
// @Summary      관리자 조회
// @Description  관리자 ID로 관리자 정보를 조회합니다
// @Tags         관리자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "관리자 ID"
// @Success      200  {object}  utils.BaseResponse "관리자 조회 성공"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      404  {object}  utils.BaseResponse "관리자 없음"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /admin/{id} [get]
func GetAdminService(c *gin.Context) {
	id := c.Param("id")
	data, err := repository.GetAdmin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}

	response := utils.CreateBaseResponse(http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// UpdateAdminService godoc
// @Summary      관리자 수정
// @Description  관리자 정보를 수정합니다. 비밀번호를 변경하지 않으려면 password 필드를 비워두세요.
// @Tags         관리자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.UpdateAdminReqDto true "관리자 수정 정보"
// @Success      200  {object}  utils.BaseResponse "관리자 수정 성공"
// @Failure      400  {object}  utils.BaseResponse "잘못된 요청"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /admin [put]
func UpdateAdminService(c *gin.Context) {
	var updateDto dto.UpdateAdminReqDto

	if err := c.ShouldBindJSON(&updateDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var hashPassword string
	if updateDto.Password != "" {
		var err error
		hashPassword, err = utils.HashPassword(updateDto.Password)
		if err != nil {
			response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	adminModel := models.Admin{
		Email:     updateDto.Email,
		Password:  hashPassword,
		AdminId:   updateDto.AdminId,
		AdminName: updateDto.AdminName,
	}

	err := repository.UpdateAdmin(adminModel)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

// DeleteAdminService godoc
// @Summary      관리자 삭제 (소프트 삭제)
// @Description  주어진 관리자 ID로 소프트 삭제를 수행합니다
// @Tags         관리자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "관리자 ID"
// @Success      200  {object}  utils.BaseResponse "관리자 삭제 성공"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      404  {object}  utils.BaseResponse "관리자 없음"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /admin/{id} [delete]
func DeleteAdminService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteAdmin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

// DeleteHardAdminService godoc
// @Summary      관리자 완전 삭제
// @Description  주어진 관리자 ID로 데이터베이스에서 완전 삭제를 수행합니다
// @Tags         관리자
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "관리자 ID"
// @Success      200  {object}  utils.BaseResponse "관리자 완전 삭제 성공"
// @Failure      401  {object}  utils.BaseResponse "인증 실패"
// @Failure      404  {object}  utils.BaseResponse "관리자 없음"
// @Failure      500  {object}  utils.BaseResponse "서버 오류"
// @Router       /admin/{id}/hard [delete]
func DeleteHardAdminService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteHardAdmin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}
