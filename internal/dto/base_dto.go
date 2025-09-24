package dto

type BaseResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BaseListResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}
type BaseListReqDto struct {
	Offset int `json:"offset" default:"0"`
	Limit  int `json:"limit" default:"10"`
}

func CreateBaseResponse(status uint, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func CreateBaseListResponse(status uint, message string, data interface{}, count int64) BaseListResponse {
	return BaseListResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Count:   count,
	}
}
