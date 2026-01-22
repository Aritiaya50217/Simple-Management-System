package domain

import "time"

type User struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Age         string    `json:"age"`
	Address     string    `json:"address"`
}

type UserRequest struct {
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address"`
}
