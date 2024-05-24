package domain

import "time"

type Report struct {
	Id          int       `json:"id" db:"id"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UserId      int       `json:"user_id" db:"user_id"`
	EquipmentId int       `json:"equipment_id" db:"equipment_id"`
}
