package usermodel

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	UpdatedAt string    `json:"updated_at" db:"updated_at"`
}

type UserRequest struct {
	Username string `json:"username" validate:"required,username_regex,min=3,max=32"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
