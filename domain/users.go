package domain

import (
	"time"

	"github.com/synt4xer/go-clean-arch/internal/dto"
)

type User struct {
	ID          uint64    `db:"id"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	FullName    string    `db:"full_name"`
	PhoneNumber string    `db:"phone_number"`
	IsActive    bool      `db:"is_active"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func CreateUser(data dto.UserCreateRequest) User {
	return User{
		Email:       data.Email,
		Password:    data.Password,
		FullName:    data.FullName,
		PhoneNumber: data.PhoneNumber,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func UpdateUser(data dto.UserUpdateRequest) User {
	return User{
		Email:       data.Email,
		Password:    data.Password,
		FullName:    data.FullName,
		PhoneNumber: data.PhoneNumber,
		IsActive:    true,
		UpdatedAt:   time.Now(),
	}
}
