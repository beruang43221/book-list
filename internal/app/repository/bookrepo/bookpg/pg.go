package bookpg

import (
	"time"

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
func (r *bookRepository) GetBookById(id uint) (*model.Book, helper.Error) {
	var book model.Book

	err := r.db.First(&book, id).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return &book, nil
}
func (r *bookRepository) UpdateBook(oldBook *model.Book, newBook *model.Book) (*model.Book, helper.Error) {
	err := r.db.Model(&oldBook).Updates(newBook)

	if err.Error != nil {
		return nil, helper.ParseError(err.Error)
	}

	return oldBook, nil
}
func (r *bookRepository) DeleteBook(book *model.Book) helper.Error {
	err := r.db.Delete(&book)

	if err.Error != nil {
		return helper.ParseError(err.Error)
	}

	return nil
}
func (r *bookRepository) GetBooksByCategoriesID(categoriesID uint) ([]model.Book, helper.Error) {
	var books []model.Book
	result := r.db.Preload("Category").Where("category_id = ?", categoriesID).Find(&books)

	if result.Error != nil {
		return nil, helper.ParseError(result.Error)
	}

	return books, nil
}
func (r *bookRepository) GetBooksByDate(startDate, endDate time.Time) ([]model.Book, helper.Error) {
	var books []model.Book

	result := r.db.Where("publication BETWEEN ? AND ?", startDate, endDate).Find(&books)

	if result.Error != nil {
		return nil, helper.ParseError(result.Error)
	}

	return books, nil
}
