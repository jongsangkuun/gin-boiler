package dto

type GetUserDto struct {
	ID int
}

type CreateUserDto struct {
	Email    string
	Username string
	Password string
}

type UpdateUserDto struct {
	ID       int
	Email    string
	Username string
	Password string
}
