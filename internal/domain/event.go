package domain

import "time"

type Event struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	Duration  time.Time `json:"duration"`
	Status    string    `json:"status"`
	UserId    int       `json:"user_id" db:"user_id"`
}
