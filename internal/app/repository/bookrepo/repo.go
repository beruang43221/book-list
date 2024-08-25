package bookrepo

import (
	"time"

	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/model"
)

type BookRepository interface {
	CreateBook(book *model.Book) (*model.Book, helper.Error)
	GetAllBooks() ([]model.Book, helper.Error)
	GetBookById(id uint) (*model.Book, helper.Error)
	UpdateBook(oldBook *model.Book, newBook *model.Book) (*model.Book, helper.Error)
	DeleteBook(book *model.Book) helper.Error
	GetBooksByCategoriesID(categoriesID uint) ([]model.Book, helper.Error)
	GetBooksByDate(startDate, endDate time.Time) ([]model.Book, helper.Error)
}
