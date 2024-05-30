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

func (mr *MaintenanceRepository) Create() {}

func (mr *MaintenanceRepository) GetById() {}

func (mr *MaintenanceRepository) GetAll() {}

func (mr *MaintenanceRepository) Delete() {}

func (mr *MaintenanceRepository) Update() {}
