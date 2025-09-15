package utils

type BaseResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateBaseResponse(status uint, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
