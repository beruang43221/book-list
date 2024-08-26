package service

import (
	"time"

	"github.com/beruang43221/book-list/internal/app/dto"
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/repository/bookrepo"
	"github.com/beruang43221/book-list/internal/app/repository/categoryrepo"
)

type BookService interface {
	CreateBook(reqBook *dto.CreateBookRequest) (*dto.CreateBookResponse, helper.Error)
	GetAllBooks() ([]dto.GetAllBooksResponse, helper.Error)
	GetBookById(id uint) (*dto.GetBookByIdResponse, helper.Error)
	UpdateBook(bookId uint, update *dto.UpdateBookRequest) (*dto.UpdateBookBookResponse, helper.Error)
	DeleteBook(bookId uint) (*dto.DeleteBookResponse, helper.Error)
	GetBooksByCategoriesID(id uint) ([]dto.GetBooksbyCategoriesIdResponse, helper.Error)
	GetBooksByDate(startDate, endDate time.Time) ([]dto.GetBooksByDateResponse, helper.Error)
	GetBooksBySearchText(title, author, publisher string) ([]dto.GetBooksBySearchTextResponse, helper.Error)
}

type bookService struct {
	bookRepository     bookrepo.BookRepository
	categoryRepository categoryrepo.CategoryRepository
}

func NewBookService(bookRepository bookrepo.BookRepository, categoryRepository categoryrepo.CategoryRepository) BookService {
	return &bookService{
		bookRepository:     bookRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *bookService) CreateBook(reqBook *dto.CreateBookRequest) (*dto.CreateBookResponse, helper.Error) {
	// Validasi categoryID dengan GetCategoryByID
	_, err := s.categoryRepository.GetCategoriesbyID(reqBook.CategoryID)
	if err != nil {
		return nil, helper.NotFound("Category ID not found")
	}

	// Konversi DTO ke Entity
	book := reqBook.ToEntity()

	// Simpan data buku ke database melalui repository
	createdBook, err := s.bookRepository.CreateBook(book)

	if err != nil {
		return nil, err
	}

	// Buat response
	response := &dto.CreateBookResponse{
		ID:          createdBook.ID,
		Title:       createdBook.Title,
		Author:      createdBook.Author,
		Publication: createdBook.Publication,
		Publisher:   createdBook.Publisher,
		Pages:       createdBook.Pages,
		CategoryID:  createdBook.CategoryID,
		CreatedAt:   createdBook.CreatedAt,
	}

	return response, nil
}
func (s *bookService) GetAllBooks() ([]dto.GetAllBooksResponse, helper.Error) {
	books, err := s.bookRepository.GetAllBooks()

	if err != nil {
		return nil, err
	}

	var response []dto.GetAllBooksResponse

	for _, book := range books {
		response = append(response, dto.GetAllBooksResponse{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Publication: book.Publication,
			Publisher:   book.Publisher,
			Pages:       book.Pages,
			CategoryID:  book.CategoryID,
			CreatedAt:   book.CreatedAt,
		})
	}

	return response, nil
}
func (s *bookService) GetBookById(id uint) (*dto.GetBookByIdResponse, helper.Error) {
	book, err := s.bookRepository.GetBookById(id)

	if err != nil {
		return nil, err
	}

	response := &dto.GetBookByIdResponse{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Publication: book.Publication,
		Publisher:   book.Publisher,
		Pages:       book.Pages,
		CategoryID:  book.CategoryID,
		CreatedAt:   book.CreatedAt,
	}

	return response, nil
}
func (s *bookService) UpdateBook(bookId uint, update *dto.UpdateBookRequest) (*dto.UpdateBookBookResponse, helper.Error) {
	oldBook, err := s.bookRepository.GetBookById(bookId)

	if err != nil {
		return nil, err
	}

	newBook := update.ToEntity()

	result, err2 := s.bookRepository.UpdateBook(oldBook, newBook)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateBookBookResponse{
		ID:          result.ID,
		Title:       result.Title,
		Author:      result.Author,
		Publication: result.Publication,
		Publisher:   result.Publisher,
		Pages:       result.Pages,
		CategoryID:  result.CategoryID,
		UpdatedAt:   result.UpdatedAt,
	}

	return response, nil
}
func (s *bookService) DeleteBook(bookId uint) (*dto.DeleteBookResponse, helper.Error) {
	book, err := s.bookRepository.GetBookById(bookId)

	if err != nil {
		return nil, err
	}

	if err := s.bookRepository.DeleteBook(book); err != nil {
		return nil, err
	}

	response := &dto.DeleteBookResponse{
		Message: "Book has been successfully deleted",
	}

	return response, nil
}
func (s *bookService) GetBooksByCategoriesID(id uint) ([]dto.GetBooksbyCategoriesIdResponse, helper.Error) {
	books, err := s.bookRepository.GetBooksByCategoriesID(id)
	if err != nil {
		return nil, err
	}

	var response []dto.GetBooksbyCategoriesIdResponse

	for _, book := range books {
		response = append(response, dto.GetBooksbyCategoriesIdResponse{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Publication: book.Publication,
			Publisher:   book.Publisher,
			Pages:       book.Pages,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
			Category: dto.Category{
				ID:        book.CategoryID,
				Name:      book.Category.Name,
				CreatedAt: book.Category.CreatedAt,
				UpdatedAt: book.Category.UpdatedAt,
			},
		})
	}
	return response, nil
}
func (s *bookService) GetBooksByDate(startDate, endDate time.Time) ([]dto.GetBooksByDateResponse, helper.Error) {
	books, err := s.bookRepository.GetBooksByDate(startDate, endDate)
	if err != nil {
		return nil, err
	}

	var responses []dto.GetBooksByDateResponse

	for _, book := range books {
		responses = append(responses, dto.GetBooksByDateResponse{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Publication: book.Publication,
			Publisher:   book.Publisher,
			Pages:       book.Pages,
			CategoryID:  book.CategoryID,
			CreatedAt:   book.CreatedAt,
		})
	}

	return responses, nil
}
func (s *bookService) GetBooksBySearchText(title, author, publisher string) ([]dto.GetBooksBySearchTextResponse, helper.Error) {
	books, err := s.bookRepository.GetBooksBySearchText(title, author, publisher)

	if err != nil {
		return nil, err
	}
	var responses []dto.GetBooksBySearchTextResponse

	for _, book := range books {
		responses = append(responses, dto.GetBooksBySearchTextResponse{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Publication: book.Publication,
			Publisher:   book.Publisher,
			Pages:       book.Pages,
			CategoryID:  book.CategoryID,
			CreatedAt:   book.CreatedAt,
		})
	}

	return responses, nil
}
