package domain

import "time"

type Event struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	Duration  []uint8   `json:"duration"`
	Status    string    `json:"status"`
	UserID    int       `json:"user_id" db:"user_id"`
	Username  *string   `json:"username" db:"username"`
}

type EventCreateDto struct {
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	StartDate    time.Time `json:"start_date" db:"start_date"`
	EndDate      time.Time `json:"end_date" db:"end_date"`
	UserID       int       `json:"user_id" db:"user_id"`
	EquipmentIDs []int     `json:"equipment_id" db:"equipment_id"`
}

type EventUpdateDto struct{}
