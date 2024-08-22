package bookrepo

import (
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/model"
)

type BookRepository interface {
	CreateBook(book *model.Book) (*model.Book, helper.Error)
	GetAllBooks() ([]model.Book, helper.Error)
}
