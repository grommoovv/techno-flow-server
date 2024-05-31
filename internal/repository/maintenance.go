package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
)

type MaintenanceRepository struct {
	db *sqlx.DB
}

func NewMaintenanceRepository(db *sqlx.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

func (mr *MaintenanceRepository) Create(dto entities.MaintenanceCreateDto) (int, error) {
	var maintenanceID int

	query := fmt.Sprintf("INSERT INTO %s () values () RETURNING id", postgres.MaintenanceTable)

	row := mr.db.QueryRow(query)
	if err := row.Scan(&maintenanceID); err != nil {
		return 0, err
	}

	return maintenanceID, nil
}

func (mr *MaintenanceRepository) GetAll() ([]entities.Maintenance, error) {
	var maintenance []entities.Maintenance

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.MaintenanceTable)

	if err := mr.db.Select(&maintenance, query); err != nil {
		return nil, err
	}

	return maintenance, nil
}

func (mr *MaintenanceRepository) GetById(id int) (entities.Maintenance, error) {
	var maintenance entities.Maintenance

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.MaintenanceTable)

	err := mr.db.QueryRow(query, id).Scan(&maintenance)

	return maintenance, err
}

func (mr *MaintenanceRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.MaintenanceTable)
	if _, err := mr.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (mr *MaintenanceRepository) Update() {}
