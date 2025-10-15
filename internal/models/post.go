package models

type Post struct {
	BaseModel
	PostData string `json:"post_data" gorm:"type:text;not null;default:''"`
	UserId   string `json:"user_id" gorm:"not null;size:255;index"`
	User     *User  `json:"user" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Post) TableName() string {
	return "posts"
}
