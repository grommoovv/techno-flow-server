package repository

import (
	"github.com/jmoiron/sqlx"
)

type MaintenanceRepository struct {
	db *sqlx.DB
}

func NewMaintenanceRepository(db *sqlx.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

func (mr *MaintenanceRepository) CreateMaintenance() {}

func (mr *MaintenanceRepository) GetMaintenance() {}

func (mr *MaintenanceRepository) GetAllMaintenance() {}

func (mr *MaintenanceRepository) DeleteMaintenance() {}

func (mr *MaintenanceRepository) UpdateMaintenance() {}
