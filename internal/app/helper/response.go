package helper

import "net/http"

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Message    *string     `json:"message,omitempty"`
}

func NewSuccessResponse(data interface{}, message *string) Response {
	return Response{
		StatusCode: http.StatusOK,
		Status:     "Success",
		Data:       data,
		Message:    message,
	}
}
