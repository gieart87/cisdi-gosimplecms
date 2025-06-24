package repositories

import (
	"gosimplecms/configs"
	"gosimplecms/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	Create(user models.User) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user models.User) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u userRepository) GetAll() ([]models.User, error) {
	var users []models.User

	result := configs.DB.Find(&users)
	return users, result.Error
}

func (u userRepository) Create(user models.User) (*models.User, error) {
	err := configs.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := configs.DB.Take(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := configs.DB.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userRepository) Update(user models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
