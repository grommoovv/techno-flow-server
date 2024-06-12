package entities

import (
	"time"
)

type User struct {
	Id          int       `json:"id" db:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	Email       *string   `json:"email"`
	FullName    *string   `json:"fullname"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Role        *string   `json:"role" db:"role"`
	AccessLevel *int      `json:"access_level" db:"access_level"`
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
