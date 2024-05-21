package repository

import (
	"github.com/jmoiron/sqlx"
)

type EquipmentRepository struct {
	db *sqlx.DB
}

func NewEquipmentRepository(db *sqlx.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (er *EquipmentRepository) CreateEquipment() {}

func (er *EquipmentRepository) GetEquipment() {}

func (er *EquipmentRepository) GetAllEquipment() {}

func (er *EquipmentRepository) DeleteEquipment() {}

func (er *EquipmentRepository) UpdateEquipment() {}
