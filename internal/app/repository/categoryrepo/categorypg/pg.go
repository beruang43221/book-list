package categorypg

import (
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/model"
	"github.com/beruang43221/book-list/internal/app/repository/categoryrepo"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) categoryrepo.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) CreateCategory(category *model.Category) (*model.Category, helper.Error) {
	if r.db == nil {
		return nil, helper.BadRequest("database connection is not initialized")
	}

	if category == nil {
		return nil, helper.BadRequest("category data is invalid")
	}
	err := r.db.Create(&category).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return category, nil
}

func (r *categoryRepository) GetAllCategories() ([]model.Category, helper.Error) {
	var categories []model.Category

	err := r.db.Find(&categories).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return categories, nil
}

func (r *categoryRepository) GetCategoriesbyID(id uint) (*model.Category, helper.Error) {
	var categories model.Category

	err := r.db.First(&categories, id).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return &categories, nil
}
