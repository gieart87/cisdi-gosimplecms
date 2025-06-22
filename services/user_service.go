package services

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"gosimplecms/models"
	"gosimplecms/repositories"
	"gosimplecms/utils/errs"
	"gosimplecms/utils/password"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	Register(registerRequest models.RegisterRequest) (*models.User, error)
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

func (s userService) Register(registerRequest models.RegisterRequest) (*models.User, error) {

	existUser, err := s.repo.FindByEmail(registerRequest.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	
	if existUser != nil {
		return nil, errs.NewAppError(errs.ErrCodeEmailAlreadyRegistered, errs.ErrMessageEmailAlreadyRegistered)
	}

	user := models.User{
		Name:       registerRequest.Name,
		Email:      registerRequest.Email,
		VerifiedAt: sql.NullTime{},
		Password:   password.HashPassword(registerRequest.Password),
		Role: func() string {
			if registerRequest.Role == "" {
				return models.RoleUser
			}
			return registerRequest.Role
		}(),
	}

	return s.repo.Create(user)
}

func (s userService) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}
