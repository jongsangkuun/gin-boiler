package middleware

import (
	"gin-boiler/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.CreateBaseResponse(http.StatusUnauthorized, "토큰이 필요합니다", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Bearer 토큰 추출
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			response := utils.CreateBaseResponse(http.StatusUnauthorized, "잘못된 토큰 형식", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		token := tokenParts[1]

		claims, err := utils.ValidateJWT(token)
		if err != nil {
			response := utils.CreateBaseResponse(http.StatusUnauthorized, "유효하지 않은 토큰", err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// 토큰이 유효하면 사용자 정보를 컨텍스트에 저장
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		c.Next()
	}
}
