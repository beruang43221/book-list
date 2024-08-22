package service

import (
	"github.com/beruang43221/book-list/internal/app/dto"
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/repository/bookrepo"
	"github.com/beruang43221/book-list/internal/app/repository/categoryrepo"
)

type BookService interface {
	CreateBook(reqBook *dto.CreateBookRequest) (*dto.CreateBookResponse, helper.Error)
	GetAllBooks() ([]dto.GetAllBooksResponse, helper.Error)
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
