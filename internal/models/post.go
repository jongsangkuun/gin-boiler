package models

type Post struct {
	BaseModel
	Title   string `json:"title" gorm:"not null;size:255"`
	Content string `json:"content" gorm:"not null;size:255"`
	UserId  string `json:"user_id" gorm:"not null;size:255"`
}

func (Post) TableName() string {
	return "posts"
}
