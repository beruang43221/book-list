package helper

import (
	"strconv"
	"time"

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

func GetQueryDateParam(context *gin.Context) (startDate, endDate string, err error) {
	startDate = context.Query("startDate")
	endDate = context.Query("endDate")
	return startDate, endDate, nil
}

func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
