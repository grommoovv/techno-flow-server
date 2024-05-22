package domain

import (
	"time"
)

type User struct {
	Id        int       `json:"id" db:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	Email     *string   `json:"email"`
	FullName  *string   `json:"fullname"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserSignInDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateDto struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	FullName *string `json:"fullname"`
}
