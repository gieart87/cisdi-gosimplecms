package repositories

import (
	"gosimplecms/configs"
	"gosimplecms/models"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	Create(category models.Category) (*models.Category, error)
	FindByID(id string) (*models.Category, error)
	Update(category models.Category) (*models.Category, error)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (c categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	err := configs.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c categoryRepository) Create(category models.Category) (*models.Category, error) {
	err := configs.DB.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c categoryRepository) FindByID(id string) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryRepository) Update(category models.Category) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}
