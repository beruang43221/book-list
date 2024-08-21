package service

import (
	"github.com/beruang43221/book-list/internal/app/dto"
	"github.com/beruang43221/book-list/internal/app/helper"
	"github.com/beruang43221/book-list/internal/app/repository/categoryrepo"
)

type CategoryService interface {
	CreateCategory(reqCategory *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, helper.Error)
	GetAllCategories() ([]dto.GetAllCategoriesResponse, helper.Error)
	GetCategoriesbyID(id uint) (*dto.GetCategoriesbyIDResponse, helper.Error)
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

func (s *categoryService) GetAllCategories() ([]dto.GetAllCategoriesResponse, helper.Error) {
	categories, err := s.categoryRepository.GetAllCategories()

	if err != nil {
		return nil, err
	}

	var response []dto.GetAllCategoriesResponse

	for _, category := range categories {
		response = append(response, dto.GetAllCategoriesResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	return response, nil
}

func (s *categoryService) GetCategoriesbyID(id uint) (*dto.GetCategoriesbyIDResponse, helper.Error) {
	categories, err := s.categoryRepository.GetCategoriesbyID(id)

	if err != nil {
		return nil, err
	}

	response := &dto.GetCategoriesbyIDResponse{
		ID:        categories.ID,
		Name:      categories.Name,
		CreatedAt: categories.CreatedAt,
		UpdatedAt: categories.UpdatedAt,
	}

	return response, nil
}
