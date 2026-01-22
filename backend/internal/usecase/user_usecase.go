package usecase

import (
	"simple-management-system/internal/domain"
	"simple-management-system/internal/repository"
	"strconv"
	"time"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetAll() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) GetByID(id int64) (*domain.User, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) Create(userReq *domain.UserRequest) error {
	age := strconv.Itoa(time.Now().Year() - userReq.DateOfBirth.Year())
	user := domain.User{
		FirstName:   userReq.FirstName,
		LastName:    userReq.LastName,
		DateOfBirth: userReq.DateOfBirth,
		Age:         age,
		Address:     userReq.Address,
	}
	return u.repo.Create(&user)
}

func (u *UserUsecase) Update(user *domain.User) error {
	return u.repo.Update(user)
}
