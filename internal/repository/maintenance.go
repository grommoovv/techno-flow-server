package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
)

var (
	ErrMaintenanceNotFound   = errors.New("maintenance not found")
	ErrMaintenanceNotCreated = errors.New("failed to create maintenance")
)

type MaintenanceRepository struct {
	db *sqlx.DB
}

func NewMaintenanceRepository(db *sqlx.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

func (mr *MaintenanceRepository) Create(ctx context.Context, dto entities.MaintenanceCreateDto) (int, error) {
	const op = "Repository/MaintenanceRepository.Create"
	var maintenanceID int

	query := fmt.Sprintf(`
		INSERT INTO %s (equipment_id, fixed_in) 
		values ($1, $2) RETURNING id`,
		postgres.MaintenanceTable)

	if err := mr.db.QueryRowContext(ctx, query, dto.EquipmentId, dto.FixedIn).Scan(&maintenanceID); err != nil {
		return 0, fmt.Errorf("%s: error during creating maintenance: %w", op, ErrMaintenanceNotCreated)
	}

	return maintenanceID, nil
}

func (mr *MaintenanceRepository) GetAll(ctx context.Context) ([]entities.Maintenance, error) {
	const op = "Repository/MaintenanceRepository.GetAll"
	var maintenance []entities.Maintenance

	query := fmt.Sprintf(`
		SELECT m.id, m.equipment_id, e.title as equipment_title, m.fixed_in, m.created_at 
		FROM %s m 
		INNER JOIN %s e on m.equipment_id = e.id 
		ORDER BY m.created_at DESC`,
		postgres.MaintenanceTable, postgres.EquipmentTable)

	if err := mr.db.SelectContext(ctx, &maintenance, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting maintenance: %w", op, ErrMaintenanceNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting maintenance: %w", op, err)
	}

	return maintenance, nil
}

func (mr *MaintenanceRepository) GetById(ctx context.Context, id int) (entities.Maintenance, error) {
	const op = "Repository/MaintenanceRepository.GetById"
	var maintenance entities.Maintenance

	query := fmt.Sprintf(`
		SELECT * 
		FROM %s 
		WHERE id = $1`,
		postgres.MaintenanceTable)

	row := mr.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&maintenance.Id, &maintenance.EquipmentId, &maintenance.FixedIn, &maintenance.CreatedAt); err != nil {
		return entities.Maintenance{}, fmt.Errorf("%s: error during getting maintenance by id: %w", op, err)
	}

	return maintenance, nil
}

func (mr *MaintenanceRepository) Delete(ctx context.Context, id int) error {
	const op = "Repository/MaintenanceRepository.Delete"
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.MaintenanceTable)
	if _, err := mr.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("%s: error during deleting maintenance by id: %w", op, err)
	}

	return nil
}

func (mr *MaintenanceRepository) Update(ctx context.Context) {}
