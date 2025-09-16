package dto

type LoginReqDto struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResDto struct {
	Token    string `json:"token"`
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
