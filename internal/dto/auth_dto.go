package dto

// LoginReqDto 로그인 요청 데이터
type LoginReqDto struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResDto 로그인 응답 데이터
type LoginResDto struct {
	Token    string `json:"token"`
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
