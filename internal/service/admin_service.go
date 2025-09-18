package service

import (
	"gin-boiler/internal/dto"
	"gin-boiler/internal/models"
	"gin-boiler/internal/repository"
	"gin-boiler/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func GetAdminService(c *gin.Context) {
	id := c.Param("id")
	data, err := repository.GetAdmin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}

	response := utils.CreateBaseResponse(http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func UpdateAdminService(c *gin.Context) {
	var updateDto dto.UpdateAdminReqDto
	var hashPassword string
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		response := utils.CreateBaseResponse(http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, response)
	}

	adminModel := models.Admin{
		Email:     updateDto.Email,
		Password:  hashPassword,
		AdminId:   updateDto.AdminId,
		AdminName: updateDto.AdminName,
	}

	if updateDto.Password != "" {
		hashPassword, err := utils.HashPassword(updateDto.Password)
		if err != nil {
			response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
			c.JSON(http.StatusInternalServerError, response)
		}
		adminModel.Password = hashPassword
	}

	err := repository.UpdateAdmin(adminModel)
	if err != nil {
		response := utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
	}

	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

func DeleteAdminService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteAdmin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

func DeleteHardAdminService(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteHardAdmin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.CreateBaseResponse(http.StatusInternalServerError, "error", err.Error()))
	}
	response := utils.CreateBaseResponse(http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}
