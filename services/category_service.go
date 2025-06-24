package services

import (
	"gosimplecms/models"
	"gosimplecms/repositories"
)

type CategoryService interface {
	Create(request models.CreateCategoryRequest) (*models.Category, error)
	Update(request models.UpdateCategoryRequest) (*models.Category, error)
	GetCategories() ([]models.Category, error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (c categoryService) Create(req models.CreateCategoryRequest) (*models.Category, error) {
	category := models.Category{
		Name: req.Name,
	}

	return c.categoryRepository.Create(category)
}

func (c categoryService) Update(request models.UpdateCategoryRequest) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryService) GetCategories() ([]models.Category, error) {
	return c.categoryRepository.GetAll()
}
