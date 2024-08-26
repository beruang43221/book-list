package dto

import (
	"time"

	"github.com/beruang43221/book-list/internal/app/model"
)

type CreateBookRequest struct {
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Publication time.Time `json:"publication" binding:"required"`
	Publisher   string    `json:"publisher" binding:"required"`
	Pages       uint      `json:"pages" binding:"required"`
	CategoryID  uint      `json:"category_id" binding:"required"`
}

func (c *CreateBookRequest) ToEntity() *model.Book {
	return &model.Book{
		Title:       c.Title,
		Author:      c.Author,
		Publication: c.Publication,
		Publisher:   c.Publisher,
		Pages:       c.Pages,
		CategoryID:  c.CategoryID,
	}
}

type CreateBookResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" `
	Author      string    `json:"author"`
	Publication time.Time `json:"publication" `
	Publisher   string    `json:"publisher" `
	Pages       uint      `json:"pages" `
	CategoryID  uint      `json:"category_id" `
	CreatedAt   time.Time `json:"created_at"`
}

type GetAllBooksResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" `
	Author      string    `json:"author"`
	Publication time.Time `json:"publication" `
	Publisher   string    `json:"publisher" `
	Pages       uint      `json:"pages" `
	CategoryID  uint      `json:"category_id" `
	CreatedAt   time.Time `json:"created_at"`
}

type GetBookByIdResponse GetAllBooksResponse

type UpdateBookRequest CreateBookRequest

func (c *UpdateBookRequest) ToEntity() *model.Book {
	return &model.Book{
		Title:       c.Title,
		Author:      c.Author,
		Publication: c.Publication,
		Publisher:   c.Publisher,
		Pages:       c.Pages,
		CategoryID:  c.CategoryID,
	}
}

type UpdateBookBookResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" `
	Author      string    `json:"author"`
	Publication time.Time `json:"publication" `
	Publisher   string    `json:"publisher" `
	Pages       uint      `json:"pages" `
	CategoryID  uint      `json:"category_id" `
	UpdatedAt   time.Time `json:"updated_at"`
}

type DeleteBookResponse struct {
	Message string `json:"message"`
}

type GetBooksbyCategoriesIdResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" `
	Author      string    `json:"author"`
	Publication time.Time `json:"publication" `
	Publisher   string    `json:"publisher" `
	Pages       uint      `json:"pages" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Category    Category  `json:"category"`
}

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBooksByDateResponse GetAllBooksResponse

type GetBooksBySearchTextResponse GetAllBooksResponse
