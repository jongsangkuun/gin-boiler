package models

type AccountStatus string

// User 사용자 모델
type User struct {
	BaseModel
	UserId        string        `json:"user_id" gorm:"uniqueIndex;not null;size:255"`
	Email         string        `json:"email" gorm:"uniqueIndex;not null;size:255"`
	Username      string        `json:"user_name" gorm:"uniqueIndex;not null;size:255"`
	Password      string        `json:"-" gorm:"not null;size:255"`
	AccountStatus AccountStatus `json:"account_status" gorm:"type:varchar(20);not null;default:pending"`
}

func (User) TableName() string {
	return "users"
}

const (
	AccountStatusPending  AccountStatus = "pending"
	AccountStatusActive   AccountStatus = "active"
	AccountStatusInactive AccountStatus = "inactive"
)
