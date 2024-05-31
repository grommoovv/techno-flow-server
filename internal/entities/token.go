package entities

type Token struct {
	Id           int    `json:"id" db:"id"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	UserId       int    `json:"user_id" db:"user_id"`
}
