package models

import (
	"gorm.io/gorm"
)

type AccountStatus string

// BaseModel Swagger를 위한 기본 모델 구조체
type BaseModel struct {
	// gorm.Model은 다음 필드들을 포함하는 기본 GoLang 구조체입니다: ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
}

const (
	AccountStatusPending  AccountStatus = "pending"
	AccountStatusActive   AccountStatus = "active"
	AccountStatusInactive AccountStatus = "inactive"
)
