package controller

import (
	"net/http"

	"github.com/beruang43221/book-list/internal/app/dto"
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	CreateBook(context *gin.Context)
	GetAllBooks(context *gin.Context)
	UpdateBook(context *gin.Context)
	DeleteBook(context *gin.Context)
	GetBooksByCategories(context *gin.Context)
	GetBooksByDate(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &bookController{
		bookService: bookService,
	}
}

func (c *bookController) CreateBook(context *gin.Context) {
	var reqBook dto.CreateBookRequest

	if err := context.ShouldBindJSON(&reqBook); err != nil {
		errorHandler := helper.UnprocessableEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := c.bookService.CreateBook(&reqBook)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, result)
}

func (c *bookController) GetAllBooks(context *gin.Context) {
	books, err := c.bookService.GetAllBooks()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, books)
}

func (c *bookController) UpdateBook(context *gin.Context) {
	bookId, _ := helper.GetIdParam(context)
	var requestBody dto.UpdateBookRequest

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		errorHandler := helper.UnprocessableEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	update, err := c.bookService.UpdateBook(bookId, &requestBody)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, update)
}

func (c *bookController) DeleteBook(context *gin.Context) {
	id, _ := helper.GetIdParam(context)

	delete, err := c.bookService.DeleteBook(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, delete)
}

func (c *bookController) GetBooksByCategories(context *gin.Context) {
	id, err := helper.GetCategoryIDParam(context)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	results, err := c.bookService.GetBooksByCategoriesID(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, results)
}

func (c *bookController) GetBooksByDate(context *gin.Context) {
	startDate, endDate, err := helper.GetQueryDateParam(context)
	if err != nil {
		errorHandler := helper.BadRequest(err.Error())
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	startDateStr, err := helper.ParseDate(startDate)
	if err != nil {
		errorHandler := helper.BadRequest(err.Error())
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	endDateStr, err := helper.ParseDate(endDate)
	if err != nil {
		errorHandler := helper.BadRequest(err.Error())
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	books, err := c.bookService.GetBooksByDate(startDateStr, endDateStr)
	if err != nil {
		errorHandler := helper.InternalServerError(err.Error())
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	context.JSON(http.StatusOK, books)
}
