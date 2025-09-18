package models

type AdminRole string

// Admin 사용자 모델
type Admin struct {
	BaseModel
	AdminId   string    `json:"admin_id" gorm:"uniqueIndex;not null;size:255"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null;size:255"`
	AdminName string    `json:"admin_name" gorm:"uniqueIndex;not null;size:255"`
	Password  string    `json:"-" gorm:"not null;size:255"`
	Role      AdminRole `json:"role" gorm:"type:varchar(20);not null;default:normal_admin"`
}

const SuperAdmin AdminRole = "super_admin"
const NormalAdmin AdminRole = "normal_admin"

func (Admin) TableName() string {
	return "admins"
}
