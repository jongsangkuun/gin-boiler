package utils

import (
	"errors"
	"gin-boiler/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(user_id, email, username string) (string, error) {
	claims := Claims{
		UserId:   user_id,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24시간
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-boiler",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// string을 []byte로 변환
	return token.SignedString([]byte(config.ENV.JwtSecret))
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// string을 []byte로 변환
		return []byte(config.ENV.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
