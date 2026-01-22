package persistence

import "time"

type UserModel struct {
	ID          int64     `gorm:"primaryKey"`
	FirstName   string    `gorm:"size:100;not null;uniqueIndex:idx_fullname"`
	LastName    string    `gorm:"size:100;not null;uniqueIndex:idx_fullname"`
	DateOfBirth time.Time `gorm:"date_of_birth"`
	Age         int64     `gorm:"age"`
	Address     string    `gorm:"address"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
