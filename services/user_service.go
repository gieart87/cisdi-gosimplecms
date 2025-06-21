package services

import (
	"gosimplecms/models"
	"gosimplecms/repositories"
	"gosimplecms/utils/password"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	CreateUser(user models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s userService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s userService) CreateUser(user models.User) (*models.User, error) {
	user.Password = password.HashPassword(user.Password)
	if user.Role == "" {
		user.Role = models.RoleUser
	}

	return s.repo.Create(user)
}

func (s userService) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}
