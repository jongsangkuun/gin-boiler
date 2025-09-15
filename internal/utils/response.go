package repository

type BaseResponse struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
