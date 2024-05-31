package entities

import "time"

type Maintenance struct {
	Id          int       `json:"id" db:"id"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
	EndDate     time.Time `json:"end_date" db:"end_date"`
	Timeline    string    `json:"timeline"`
	EquipmentId int       `json:"equipment_id" db:"equipment_id"`
}

type MaintenanceCreateDto struct {
}
