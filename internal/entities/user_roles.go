package entities

import "time"

type UserRoles struct {
	Id          int       `json:"id" db:"id"`
	UserId      int       `json:"user_id" db:"user_id"`
	Role        string    `json:"role" db:"role"`
	AccessLevel int       `json:"access_level" db:"access_level"`
	ModifiedAt  time.Time `json:"modified_at" db:"modified_at"`
}
