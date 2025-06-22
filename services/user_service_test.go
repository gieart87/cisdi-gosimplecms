package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gosimplecms/models"
	"gosimplecms/repositories"
	"gosimplecms/utils/errs"
	"gosimplecms/utils/password"
	"testing"
)

func TestNewUserService(t *testing.T) {
	mockRepo := repositories.NewUserRepositoryMock()
	service := NewUserService(mockRepo)

	assert.NotNil(t, service)
}

func TestUserService_RegisterSuccess(t *testing.T) {
	mockRepo := repositories.NewUserRepositoryMock()

	mockRepo.FindByEmailFunc = func(email string) (*models.User, error) {
		return nil, gorm.ErrRecordNotFound
	}

	mockRepo.CreateFunc = func(user models.User) (*models.User, error) {
		user.ID = uuid.New().String()
		return &user, nil
	}

	service := NewUserService(mockRepo)

	req := models.RegisterRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret123",
	}

	user, err := service.Register(req)

	assert.Nil(t, err)
	assert.Equal(t, req.Email, user.Email)
	assert.Equal(t, req.Name, user.Name)
}

func TestUserService_RegisterEmailAlreadyExist(t *testing.T) {
	mockRepo := repositories.NewUserRepositoryMock()

	existingUser := models.User{
		ID:    uuid.New().String(),
		Name:  "Alice",
		Email: "alice@example.com",
	}

	mockRepo.FindByEmailFunc = func(email string) (*models.User, error) {
		return &existingUser, nil
	}

	service := NewUserService(mockRepo)

	req := models.RegisterRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret123",
	}

	_, err := service.Register(req)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("%s: %s", errs.ErrCodeEmailAlreadyRegistered, errs.ErrMessageEmailAlreadyRegistered), err.Error())
}

func TestUserService_LoginSuccess(t *testing.T) {
	mockRepo := repositories.NewUserRepositoryMock()
	plainPassword := "secret123"
	hashedPassword := password.HashPassword(plainPassword)

	mockRepo.FindByEmailFunc = func(id string) (*models.User, error) {
		return &models.User{
			ID:       uuid.New().String(),
			Name:     "John",
			Email:    "john@example.com",
			Password: hashedPassword,
			Role:     models.RoleUser,
		}, nil
	}

	service := NewUserService(mockRepo)
	token, err := service.Login(models.LoginRequest{
		Email:    "john@example.com",
		Password: plainPassword,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestUserService_LoginUserNotFound(t *testing.T) {
	mockRepo := repositories.NewUserRepositoryMock()

	mockRepo.FindByEmailFunc = func(id string) (*models.User, error) {
		return &models.User{}, nil
	}

	service := NewUserService(mockRepo)
	_, err := service.Login(models.LoginRequest{
		Email:    "invalid@example.com",
		Password: "password123",
	})

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("%s: %s", errs.ErrCodeLoginFailed, errs.ErrMessageLoginFailed), err.Error())
}

func TestUserService_LoginInvalidPassword(t *testing.T) {
	mockRepo := repositories.NewUserRepositoryMock()
	plainPassword := "CorrectPassword"
	hashedPassword := password.HashPassword(plainPassword)

	mockRepo.FindByEmailFunc = func(id string) (*models.User, error) {
		return &models.User{
			ID:       uuid.New().String(),
			Name:     "John",
			Email:    "john@example.com",
			Password: hashedPassword,
			Role:     models.RoleUser,
		}, nil
	}

	service := NewUserService(mockRepo)
	_, err := service.Login(models.LoginRequest{
		Email:    "john@example.com",
		Password: "WrongPassword",
	})

	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("%s: %s", errs.ErrCodeLoginFailed, errs.ErrMessageLoginFailed), err.Error())
}
