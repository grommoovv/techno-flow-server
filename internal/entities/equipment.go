package entities

import "time"

type Equipment struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	IsAvailable *bool     `json:"is_available" db:"is_available"`
}

type EquipmentCreateDto struct {
	Title  string `json:"title" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type EquipmentUpdateDto struct {
	Title  *string `json:"title"`
	Status *string `json:"status"`
	//IsAvailable *bool      `json:"is_available" db:"is_available"`
	//AvailableAt *time.Time `json:"available_at" db:"available_at"`
	//ReservedAt  *time.Time `json:"reserved_at" db:"reserved_at"`
	//UserId      *int       `json:"user_id" db:"user_id"`
}

type EquipmentUsageHistory struct {
	ID         int       `json:"id" db:"id"`
	Username   string    `json:"username" db:"username"`
	EventTitle string    `json:"event_title" db:"event_title"`
	StartDate  time.Time `json:"start_date" db:"start_date"`
	EndDate    time.Time `json:"end_date" db:"end_date"`
}

type GetAvailableEquipmentByDateDto struct {
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
}
