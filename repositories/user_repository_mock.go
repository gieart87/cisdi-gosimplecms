package repositories

import "gosimplecms/models"

type UserRepositoryMock struct {
	Users           []models.User
	CreateFunc      func(user models.User) (*models.User, error)
	FindByIDFunc    func(id string) (*models.User, error)
	FindByEmailFunc func(email string) (*models.User, error)
	UpdateFunc      func(user models.User) (*models.User, error)
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		Users: []models.User{},
		CreateFunc: func(user models.User) (*models.User, error) {
			return &user, nil
		},
		FindByIDFunc: func(id string) (*models.User, error) {
			return &models.User{}, nil
		},
		FindByEmailFunc: func(email string) (*models.User, error) {
			return &models.User{}, nil
		},
		UpdateFunc: func(user models.User) (*models.User, error) {
			return &user, nil
		},
	}
}

func (u *UserRepositoryMock) GetAll() ([]models.User, error) {
	return u.Users, nil
}

func (u *UserRepositoryMock) Create(user models.User) (*models.User, error) {
	return u.CreateFunc(user)
}

func (u *UserRepositoryMock) FindByID(id string) (*models.User, error) {
	return u.FindByIDFunc(id)
}

func (u *UserRepositoryMock) FindByEmail(email string) (*models.User, error) {
	return u.FindByEmailFunc(email)
}

func (u *UserRepositoryMock) Update(user models.User) (*models.User, error) {
	return u.UpdateFunc(user)
}
