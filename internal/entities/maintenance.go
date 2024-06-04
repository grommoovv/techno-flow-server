package entities

import "time"

type Maintenance struct {
	Id             int       `json:"id" db:"id"`
	EquipmentId    int       `json:"equipment_id" db:"equipment_id"`
	EquipmentTitle string    `json:"equipment_title" db:"equipment_title"`
	FixedIn        time.Time `json:"fixed_in" db:"fixed_in"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type MaintenanceCreateDto struct {
	EquipmentId int       `json:"equipment_id" db:"equipment_id"`
	FixedIn     time.Time `json:"fixed_in" db:"fixed_in"`
}
