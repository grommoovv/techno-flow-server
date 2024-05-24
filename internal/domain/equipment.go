package domain

import "time"

type Equipment struct {
	Id          int        `json:"id" db:"id"`
	Title       string     `json:"title"`
	State       string     `json:"state"`
	IsAvailable bool       `json:"is_available" db:"is_available"`
	AvailableAt *time.Time `json:"available_at " db:"available_at"`
	ReservedAt  *time.Time `json:"reserved_at" db:"reserved_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UserId      *int       `json:"user_id" db:"user_id"`
}

type EquipmentCreateDto struct {
	Title string `json:"title" binding:"required"`
	State string `json:"state" binding:"required"`
}

type EquipmentUpdateDto struct {
	Title       *string    `json:"title"`
	State       *string    `json:"state"`
	IsAvailable *bool      `json:"is_available"`
	AvailableAt *time.Time `json:"available_at" db:"available_at"`
	ReservedAt  *string    `json:"reserved_at"`
	UserId      *int       `json:"user_id" db:"user_id"`
}

type EquipmentUsage struct {
	Id          int `json:"id" db:"id"`
	UserId      int `json:"user_id" db:"user_id"`
	EventId     int `json:"event_id" db:"event_id"`
	EquipmentId int `json:"equipment_id" db:"equipment_id"`
}
