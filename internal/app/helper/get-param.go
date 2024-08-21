package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdParam(context *gin.Context) (uint, Error) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return uint(0), BadRequest("ID param must be an integer")
	}

	return uint(id), nil
}

func GetCategoryIDParam(context *gin.Context) (uint, Error) {
	categoryIdParam := context.Param("category_id")
	if categoryIdParam == "" {
		return (0), BadRequest("Category ID parameter is required")
	}

	categoryID, err := strconv.Atoi(categoryIdParam)

	if err != nil {
		return (0), BadRequest("Category ID parameter must be an integer")
	}

	if categoryID <= 0 {
		return 0, BadRequest("Category ID must be a positive integer")
	}

	return uint(categoryID), nil
}

// func GetStringTitle(context *gin.Context, title string) (string, Error) {
// 	titleValue := context.Query(title)
// 	if titleValue == "" {
// 		return "", BadRequest("Parameter is required")

// 	}
// 	return titleValue, nil
// }

// func GetStringAuthor(context *gin.Context, paramName string) (string, Error) {
// 	paramValue := context.Param(paramName)
// 	if paramValue == "" {
// 		return "", BadRequest("Parameter is required")

// 	}
// 	return paramValue, nil
// }
// func GetStringPublisher(context *gin.Context, paramName string) (string, Error) {
// 	paramValue := context.Param(paramName)
// 	if paramValue == "" {
// 		return "", BadRequest("Parameter is required")

// 	}
// 	return paramValue, nil
// }
