package helper

import "net/http"

type Error interface {
	Error() string
	Status() int
	Type() string
}

type errorResponse struct {
	message string
	status  int
	errType string
}

func (e *errorResponse) Error() string {
	return e.message
}

func (e *errorResponse) Status() int {
	return e.status
}

func (e *errorResponse) Type() string {
	return e.errType
}

func NewError(message string, status int, errType string) Error {
	return &errorResponse{
		message: message,
		status:  status,
		errType: errType,
	}
}

func BadRequest(message string) Error {
	return NewError(message, http.StatusBadRequest, "Bad Request")
}

func Unauthorized(message string) Error {
	return NewError(message, http.StatusUnauthorized, "Unauthorized")
}

func NotFound(message string) Error {
	return NewError(message, http.StatusNotFound, "Not Found")
}

func UnprocessableEntity(message string) Error {
	return NewError(message, http.StatusUnprocessableEntity, "Invalid Request")
}

func InternalServerError(message string) Error {
	return NewError(message, http.StatusInternalServerError, "Server Error")
}
