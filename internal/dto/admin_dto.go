package dto

type CreateAdminReqDto struct {
	Email     string `json:"email"`
	AdminId   string `json:"admin_id"`
	AdminName string `json:"admin_name"`
	Password  string `json:"password"`
}

type UpdateAdminReqDto struct {
	ID uint `json:"id"`
	CreateAdminReqDto
}
