package entities

import "time"

type EquipmentUsage struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	EventID     int       `json:"event_id" db:"event_id"`
	EquipmentID int       `json:"equipment_id" db:"equipment_id"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
	EndDate     time.Time `json:"end_date" db:"end_date"`
}
