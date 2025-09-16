package dto

// CreateUserReqDto 사용자 생성 요청 데이터
type CreateUserReqDto struct {
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdateUserReqDto 사용자 수정 요청 데이터
type UpdateUserReqDto struct {
	ID int
	CreateUserReqDto
}

// ListUserResDto 사용자 목록 응답 데이터
type ListUserResDto struct {
	ID       int
	Email    string
	UserId   string
	Username string
}
