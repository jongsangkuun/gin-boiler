package dto

type BaseResponseDto struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BaseListResponseDto struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}
type BasePageReqDto struct {
	Offset int `json:"offset" default:"0"`
	Limit  int `json:"limit" default:"10"`
}

func CreateBaseResponse(status uint, message string, data interface{}) BaseResponseDto {
	return BaseResponseDto{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func CreateBaseListResponse(status uint, message string, data interface{}, count int64) BaseListResponseDto {
	return BaseListResponseDto{
		Status:  status,
		Message: message,
		Data:    data,
		Count:   count,
	}
}
