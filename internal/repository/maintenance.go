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

	query := fmt.Sprintf("INSERT INTO %s (equipment_id, fixed_in) values ($1, $2) RETURNING id", postgres.MaintenanceTable)

	row := mr.db.QueryRow(query, dto.EquipmentId, dto.FixedIn)
	if err := row.Scan(&maintenanceID); err != nil {
		return 0, err
	}

	return maintenanceID, nil
}

func (mr *MaintenanceRepository) GetAll() ([]entities.Maintenance, error) {
	var maintenance []entities.Maintenance

	query := fmt.Sprintf(`
						SELECT m.id, m.equipment_id, e.title as equipment_title, m.fixed_in, m.created_at 
						FROM %s m 
						INNER JOIN %s e on m.equipment_id = e.id 
						ORDER BY m.created_at DESC`,
		postgres.MaintenanceTable, postgres.EquipmentTable)

	if err := mr.db.Select(&maintenance, query); err != nil {
		return nil, err
	}

	return maintenance, nil
}

func (mr *MaintenanceRepository) GetById(id int) (entities.Maintenance, error) {
	var maintenance entities.Maintenance

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.MaintenanceTable)

	err := mr.db.QueryRow(query, id).Scan(&maintenance.Id, &maintenance.EquipmentId, &maintenance.FixedIn, &maintenance.CreatedAt)

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
