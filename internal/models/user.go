package models

import "fmt"
import "encoding/json"

type User struct {
	BaseModel
	UserId        string        `json:"user_id" gorm:"uniqueIndex;not null;size:255"`
	Email         string        `json:"email" gorm:"uniqueIndex;not null;size:255"`
	Username      string        `json:"username" gorm:"uniqueIndex;not null;size:255"`
	Password      string        `json:"-" gorm:"not null;size:255"`
	AccountStatus AccountStatus `json:"account_status" gorm:"type:varchar(20);not null;default:pending"`
}

func (User) TableName() string {
	return "users"
}

func (as *AccountStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	status := AccountStatus(s)
	if !status.IsValid() {
		return fmt.Errorf("invalid account status: %s", s)
	}

	*as = status
	return nil
}

func (as AccountStatus) IsValid() bool {
	switch as {
	case AccountStatusPending, AccountStatusActive, AccountStatusInactive:
		return true
	}
	return false
}
