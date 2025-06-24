package services

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"gosimplecms/models"
	"gosimplecms/repositories"
	"gosimplecms/utils/errs"
	"gosimplecms/utils/jwt"
	"gosimplecms/utils/password"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	Register(registerRequest models.RegisterRequest) (*models.User, error)
	Login(loginRequest models.LoginRequest) (string, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
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
		Role:       models.RoleUser,
	}

	return s.repo.Create(user)
}

func (s userService) Login(loginRequest models.LoginRequest) (string, error) {
	existUser, err := s.repo.FindByEmail(loginRequest.Email)
	if err != nil || !password.CheckPassword(existUser.Password, loginRequest.Password) {
		return "", errs.NewAppError(errs.ErrCodeLoginFailed, errs.ErrMessageLoginFailed)
	}

	token, err := jwt.GenerateToken(existUser.ID, existUser.Role)
	if err != nil {
		return "", errs.NewAppError(errs.ErrCodeLoginFailed, errs.ErrMessageLoginFailed)
	}

	return token, nil
}

func (s userService) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s userService) FindByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}
