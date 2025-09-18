package dto

// UserLoginReqDto 로그인 요청 데이터
type UserLoginReqDto struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserLoginResDto 로그인 응답 데이터
type UserLoginResDto struct {
	Token    string `json:"token"`
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type AdminLoginReqDto struct {
	AdminId  string `json:"admin_id" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type AdminLoginResDto struct {
	Token     string `json:"token"`
	AdminId   string `json:"admin_id"`
	Email     string `json:"email"`
	AdminName string `json:"admin_name"`
}
