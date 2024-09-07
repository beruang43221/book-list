package helper

import (
	"net/http"
	"strings"
)

func ParseError(err error) Error {
	if strings.Contains(err.Error(), "record not found") {
		return &customError{
			message: "Data not found",
			status:  http.StatusNotFound,
			errType: "Not Found",
		}
	}

	return InternalServerError("Something went wrong: " + err.Error())
}

func BadRequest(message string) Error {
	return &customError{
		message: message,
		status:  http.StatusBadRequest,
		errType: "Bad Request",
	}
}

func NotFound(message string) Error {
	return &customError{
		message: message,
		status:  http.StatusNotFound,
		errType: "Not Found",
	}
}

func InternalServerError(message string) Error {
	return &customError{
		message: message,
		status:  http.StatusInternalServerError,
		errType: "Internal Server Error",
	}
}
