package categoryrepo

import (
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/model"
)

type CategoryRepository interface {
	CreateCategory(category *model.Category) (*model.Category, helper.Error)
	GetAllCategories() ([]model.Category, helper.Error)
	GetCategoriesbyID(id uint) (*model.Category, helper.Error)
}
