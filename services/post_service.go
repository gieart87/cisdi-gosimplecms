package services

import (
	"gosimplecms/models"
	"gosimplecms/repositories"
)

type PostService interface {
	GetPosts() ([]models.Post, error)
	FindBySlug(string) (*models.Post, error)
}

type postService struct {
	repo repositories.PostRepository
}

func (p postService) FindBySlug(s string) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo: repo}
}

func (p postService) GetPosts() ([]models.Post, error) {
	return nil, nil
}
