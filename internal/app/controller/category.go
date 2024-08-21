package controller

import (
	"net/http"

	"github.com/beruang43221/book-list/internal/app/dto"
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/service"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory(context *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (c *categoryController) CreateCategory(context *gin.Context) {
	var reqCategory dto.CreateCategoryRequest

	if err := context.ShouldBindJSON(&reqCategory); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := c.categoryService.CreateCategory(&reqCategory)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, result)
}
