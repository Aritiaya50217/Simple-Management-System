package persistence

import (
	"simple-management-system/internal/domain"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	age, err := strconv.ParseInt(user.Age, 10, 64)
	if err != nil {
		return err
	}

	model := UserModel{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateOfBirth: user.DateOfBirth,
		Age:         age,
		Address:     user.Address,
	}

	if err := r.db.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(user *domain.User) error {
	var model UserModel
	return r.db.Model(&model).Where("id = ?", user.ID).Updates(model).Error
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
	var models []UserModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, 0)
	for _, m := range models {
		age := strconv.Itoa(int(m.Age))
		users = append(users, domain.User{
			ID:          m.ID,
			FirstName:   m.FirstName,
			LastName:    m.LastName,
			DateOfBirth: m.DateOfBirth,
			Age:         age + " ปี",
			Address:     m.Address,
		})
	}
	return users, nil
}

func (r *UserRepository) FindByID(id int64) (*domain.User, error) {
	var model UserModel
	err := r.db.First(&model, id).Error
	if err != nil {
		return nil, err
	}
	
	age := strconv.Itoa(int(time.Now().Year() - model.DateOfBirth.Year()))

	return &domain.User{
		ID:          model.ID,
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		DateOfBirth: model.DateOfBirth,
		Age:         age + " ปี",
		Address:     model.Address,
	}, nil
}
