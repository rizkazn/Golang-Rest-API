package models

import "time"

type User struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,gte=6"`
	Role      string    `json:"role" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

func CreateUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
