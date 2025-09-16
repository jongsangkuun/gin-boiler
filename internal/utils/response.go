package utils

type BaseResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BaseListResponse[T any] struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Data    []T    `json:"data"`
	Count   int64  `json:"count"`
}

func CreateBaseResponse(status uint, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func CreateBaseListResponse[T any](status uint, message string, data []T, count int64) BaseListResponse[T] {
	return BaseListResponse[T]{
		Status:  status,
		Message: message,
		Data:    data,
		Count:   count,
	}
}
