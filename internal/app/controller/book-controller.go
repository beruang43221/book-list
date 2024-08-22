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
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

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
