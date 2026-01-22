package repository

import "simple-management-system/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	FindAll() ([]domain.User, error)
	FindByID(id int64) (*domain.User, error)
}
