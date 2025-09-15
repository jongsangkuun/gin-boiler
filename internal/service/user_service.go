package service

import "github.com/gin-gonic/gin"

func UserPingService(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateUserService(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}
func GetUserService(c *gin.Context)        {}
func UpdateUserService(c *gin.Context)     {}
func DeleteUserService(c *gin.Context)     {}
func ListUserService(c *gin.Context)       {}
func DeleteHardUserService(c *gin.Context) {}
