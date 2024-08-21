package service

import (
	"github.com/beruang43221/book-list/internal/app/dto"
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/repository/categoryrepo"
)

type CategoryService interface {
	CreateCategory(reqCategory *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, helper.Error)
}

type categoryService struct {
	categoryRepository categoryrepo.CategoryRepository
}

func NewCategoryService(categoryRepository categoryrepo.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}

}

func (s *categoryService) CreateCategory(reqCategory *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, helper.Error) {
	category := reqCategory.ToEntity()

	createdCategory, err := s.categoryRepository.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	response := &dto.CreateCategoryResponse{
		ID:        createdCategory.ID,
		Name:      createdCategory.Name,
		CreatedAt: createdCategory.CreatedAt,
	}

	return response, nil
}
