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
	GetAllCategories(context *gin.Context)
	GetCategoriesbyID(context *gin.Context)
	UpdateCategory(context *gin.Context)
	DeleteCategory(context *gin.Context)
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
		errorHandler := helper.UnprocessableEntity("Invalid JSON body")

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

func (c *categoryController) GetAllCategories(context *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, categories)
}

func (c *categoryController) GetCategoriesbyID(context *gin.Context) {
	id, err := helper.GetIdParam(context)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	categories, err := c.categoryService.GetCategoriesbyID(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, categories)
}

func (c *categoryController) UpdateCategory(context *gin.Context) {
	categoryId, _ := helper.GetIdParam(context)
	var requestBody dto.UpdateCategoryRequest

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		errorHandler := helper.UnprocessableEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	update, err := c.categoryService.UpdateCategory(categoryId, &requestBody)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, update)
}

func (c *categoryController) DeleteCategory(context *gin.Context) {
	id, _ := helper.GetIdParam(context)

	delete, err := c.categoryService.DeleteCategory(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, delete)
}
