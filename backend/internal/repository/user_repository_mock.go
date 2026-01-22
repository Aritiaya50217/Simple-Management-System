package repository

import "simple-management-system/internal/domain"

type UserRepositoryMock struct {
	FindAllFn  func() ([]domain.User, error)
	FindByIDFn func(id int64) (*domain.User, error)
	CreateFn   func(user *domain.User) error
	UpdateFn   func(user *domain.User) error
}

func (m *UserRepositoryMock) FindAll() ([]domain.User, error) {
	return m.FindAllFn()
}

func (m *UserRepositoryMock) FindByID(id int64) (*domain.User, error) {
	return m.FindByIDFn(id)
}

func (m *UserRepositoryMock) Create(user *domain.User) error {
	return m.CreateFn(user)
}

func (m *UserRepositoryMock) Update(user *domain.User) error {
	return m.UpdateFn(user)
}
