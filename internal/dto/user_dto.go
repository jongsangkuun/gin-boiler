package dto

type GetUserDto struct {
	ID int
}

type CreateUserDto struct {
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserDto struct {
	ID       int
	Email    string
	Username string
	Password string
}
