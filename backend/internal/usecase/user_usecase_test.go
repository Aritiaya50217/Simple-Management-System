package usecase

import (
	"errors"
	"simple-management-system/internal/domain"
	"simple-management-system/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		CreateFn: func(user *domain.User) error {
			assert.Equal(t, "John", user.FirstName)
			assert.NotEmpty(t, user.Age)
			return nil
		},
	}

	uc := NewUserUsecase(mockRepo)

	err := uc.Create(&domain.UserRequest{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
		Address:     "Bangkok",
	})

	assert.NoError(t, err)
}

func TestCreateUser_Failed(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		CreateFn: func(user *domain.User) error {
			return errors.New("db error")
		},
	}

	uc := NewUserUsecase(mockRepo)

	err := uc.Create(&domain.UserRequest{})

	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		FindByIDFn: func(id int64) (*domain.User, error) {
			return &domain.User{
				ID:        id,
				FirstName: "User",
			}, nil
		},
	}

	uc := NewUserUsecase(mockRepo)

	user, err := uc.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(1), user.ID)
}

func TestGetByID_NotFound(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		FindByIDFn: func(id int64) (*domain.User, error) {
			return nil, errors.New("not found")
		},
	}

	uc := NewUserUsecase(mockRepo)

	user, err := uc.GetByID(99)

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.EqualError(t, err, "not found")
}

func TestGetAll_Success(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		FindAllFn: func() ([]domain.User, error) {
			return []domain.User{
				{ID: 1, FirstName: "A"},
				{ID: 2, FirstName: "B"},
			}, nil
		},
	}

	uc := NewUserUsecase(mockRepo)

	users, err := uc.GetAll()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestGetAll_Empty(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		FindAllFn: func() ([]domain.User, error) {
			return []domain.User{}, nil
		},
	}

	uc := NewUserUsecase(mockRepo)

	users, err := uc.GetAll()

	assert.NoError(t, err)
	assert.Empty(t, users)
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		UpdateFn: func(user *domain.User) error {
			assert.Equal(t, int64(1), user.ID)
			return nil
		},
	}

	uc := NewUserUsecase(mockRepo)

	err := uc.Update(&domain.User{
		ID:        1,
		FirstName: "Updated",
	})

	assert.NoError(t, err)
}

func TestUpdateUser_Failed(t *testing.T) {
	mockRepo := &repository.UserRepositoryMock{
		UpdateFn: func(user *domain.User) error {
			return errors.New("update failed")
		},
	}

	uc := NewUserUsecase(mockRepo)

	err := uc.Update(&domain.User{ID: 1})

	assert.Error(t, err)
	assert.EqualError(t, err, "update failed")
}

