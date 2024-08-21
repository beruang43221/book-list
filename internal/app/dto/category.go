package dto

import (
	"time"

	"github.com/beruang43221/book-list/internal/app/model"
)

type CreateCategoryRequest struct {
	Name string `json:"Name" binding:"required"`
}

func (c *CreateCategoryRequest) ToEntity() *model.Category {
	return &model.Category{
		Name: c.Name,
	}
}

type CreateCategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
