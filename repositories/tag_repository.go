package repositories

import (
	"gosimplecms/configs"
	"gosimplecms/models"
)

type TagRepository interface {
	FirstOrCreate(tag models.Tag) (*models.Tag, error)
	FindByIDs(ids []uint) ([]models.Tag, error)
}

type tagRepository struct{}

func (t tagRepository) FindByIDs(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag

	err := configs.DB.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func NewTagRepository() TagRepository {
	return &tagRepository{}
}

func (t tagRepository) FirstOrCreate(req models.Tag) (*models.Tag, error) {
	var tag models.Tag

	err := configs.DB.FirstOrCreate(&tag, req).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
