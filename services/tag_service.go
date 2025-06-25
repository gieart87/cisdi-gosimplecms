package services

import (
	"gosimplecms/models"
	"gosimplecms/repositories"
)

type TagService interface {
	Create(request models.CreateTagRequest) (*models.Tag, error)
	GetTags() ([]models.Tag, error)
}

type tagService struct {
	tagRepository repositories.TagRepository
}

func NewTagService(tagRepository repositories.TagRepository) TagService {
	return &tagService{tagRepository: tagRepository}
}

func (c tagService) Create(req models.CreateTagRequest) (*models.Tag, error) {
	category := models.Tag{
		Name: req.Name,
	}

	return c.tagRepository.Create(category)
}

func (c tagService) GetTags() ([]models.Tag, error) {
	return c.tagRepository.GetTags()
}
