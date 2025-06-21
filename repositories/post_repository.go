package repositories

import "gosimplecms/models"

type PostRepository interface {
	GetAll() ([]models.Post, error)
	Create(post models.Post) (models.Post, error)
	FindByID(id string) (models.Post, error)
	FindBySlug(slug string) (models.Post, error)
	Update(post models.Post) (models.Post, error)
}

type postRepository struct{}

func NewPostRepository() PostRepository {
	return &postRepository{}
}

func (p postRepository) GetAll() ([]models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) Create(post models.Post) (models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) FindByID(id string) (models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) FindBySlug(slug string) (models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) Update(post models.Post) (models.Post, error) {
	//TODO implement me
	panic("implement me")
}
