package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware CORS 설정을 위한 미들웨어
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 허용할 도메인 설정 (개발 환경에서는 모든 도메인 허용)
		c.Header("Access-Control-Allow-Origin", "*")

		// 허용할 HTTP 메서드 설정
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")

		// 허용할 헤더 설정
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		// 인증 정보 포함 허용
		c.Header("Access-Control-Allow-Credentials", "true")

		// Preflight 요청 처리
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// CORSMiddlewareWithConfig 설정 가능한 CORS 미들웨어
func CORSMiddlewareWithConfig(allowOrigins []string, allowMethods []string, allowHeaders []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Origin 검증
		if len(allowOrigins) > 0 {
			allowed := false
			for _, allowedOrigin := range allowOrigins {
				if allowedOrigin == "*" || allowedOrigin == origin {
					allowed = true
					break
				}
			}
			if allowed {
				c.Header("Access-Control-Allow-Origin", origin)
			}
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

		// 메서드 설정
		if len(allowMethods) > 0 {
			methods := ""
			for i, method := range allowMethods {
				if i > 0 {
					methods += ", "
				}
				methods += method
			}
			c.Header("Access-Control-Allow-Methods", methods)
		} else {
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		}

		// 헤더 설정
		if len(allowHeaders) > 0 {
			headers := ""
			for i, header := range allowHeaders {
				if i > 0 {
					headers += ", "
				}
				headers += header
			}
			c.Header("Access-Control-Allow-Headers", headers)
		} else {
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		}

		c.Header("Access-Control-Allow-Credentials", "true")

		// Preflight 요청 처리
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// ProductionCORSMiddleware 프로덕션 환경용 CORS 미들웨어
func ProductionCORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// 허용된 도메인인지 확인
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400") // 24시간 캐시
		}

		// Preflight 요청 처리
		if c.Request.Method == "OPTIONS" {
			if allowed {
				c.Status(http.StatusNoContent)
			} else {
				c.Status(http.StatusForbidden)
			}
			c.Abort()
			return
		}

		// 허용되지 않은 Origin에서의 요청 차단
		if !allowed && origin != "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
