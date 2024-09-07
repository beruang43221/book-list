package helper

import "net/http"

type Error interface {
	Error() string
	Status() int
	Type() string
}

type customError struct {
	message string
	status  int
	errType string
}

func (e *customError) Error() string {
	return e.message
}

func (e *customError) Status() int {
	return e.status
}

func (e *customError) Type() string {
	return e.errType
}

func NewErrorResponse(message string, data interface{}, code int) Response {
	return Response{
		StatusCode: code,
		Status:     "error",
		Data:       data,
		Message:    &message,
	}
}

func Unauthorized(message string) Response {
	return NewErrorResponse(message, nil, http.StatusUnauthorized)
}

func UnprocessableEntity(message string) Response {
	return NewErrorResponse(message, nil, http.StatusUnprocessableEntity)
}
