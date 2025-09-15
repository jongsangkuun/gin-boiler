package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 평문 비밀번호를 bcrypt로 해싱합니다
func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("비밀번호는 비어있을 수 없습니다")
	}

	// bcrypt를 사용하여 비밀번호 해싱 (cost 12 사용)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("비밀번호 해싱 실패: %v", err)
		return "", errors.New("비밀번호 해싱에 실패했습니다")
	}

	return string(hashedBytes), nil
}

// VerifyPassword 평문 비밀번호와 해시된 비밀번호를 비교합니다
func VerifyPassword(plainPassword, hashedPassword string) error {
	if len(plainPassword) == 0 {
		return errors.New("비밀번호는 비어있을 수 없습니다")
	}

	if len(hashedPassword) == 0 {
		return errors.New("해시된 비밀번호는 비어있을 수 없습니다")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("비밀번호가 일치하지 않습니다")
		}
		log.Printf("비밀번호 검증 중 오류 발생: %v", err)
		return errors.New("비밀번호 검증에 실패했습니다")
	}

	return nil
}
