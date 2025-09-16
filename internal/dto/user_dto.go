package dto

type CreateUserReqDto struct {
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserReqDto struct {
	ID int
	CreateUserReqDto
}

type ListUserResDto struct {
	ID       int
	Email    string
	UserId   string
	Username string
}
