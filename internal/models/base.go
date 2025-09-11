package models

import (
	"gorm.io/gorm"
)

type AccountStatus string
type BaseModel struct {
	// gorm.Model은 다음 필드들을 포함하는 기본 GoLang 구조체입니다: ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
}

const (
	AccountStatusPending  AccountStatus = "pending"
	AccountStatusActive   AccountStatus = "active"
	AccountStatusInactive AccountStatus = "inactive"
)
