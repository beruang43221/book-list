package bookpg

import (
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/model"
	"github.com/beruang43221/book-list/internal/app/repository/bookrepo"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) bookrepo.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) CreateBook(book *model.Book) (*model.Book, helper.Error) {

	err := r.db.Create(&book).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return book, nil
}
func (r *bookRepository) GetAllBooks() ([]model.Book, helper.Error) {
	var books []model.Book

	err := r.db.Find(&books).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return books, nil
}
