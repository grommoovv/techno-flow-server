package entities

type Role struct {
	Id          int    `json:"id" db:"id"`
	Role        string `json:"role" db:"role"`
	AccessLevel int    `json:"access_level" db:"access_level"`
	Description string `json:"description" db:"description"`
}
