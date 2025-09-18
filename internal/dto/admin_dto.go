package dto

// CreateAdminReqDto 관리자 생성 요청 데이터
type CreateAdminReqDto struct {
	Email     string `json:"email"`
	AdminId   string `json:"admin_id"`
	AdminName string `json:"admin_name"`
	Password  string `json:"password"`
}

// UpdateAdminReqDto 관리자 수정 요청 데이터
type UpdateAdminReqDto struct {
	ID uint `json:"id"`
	CreateAdminReqDto
}
