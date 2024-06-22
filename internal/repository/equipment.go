package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	ErrEquipmentNotFound    = errors.New("equipment not found")
	ErrUsageHistoryNotFound = errors.New("usage history not found")
	ErrNotUpdated           = errors.New("failed to update equipment")
)

type EquipmentRepository struct {
	db *sqlx.DB
}

func NewEquipmentRepository(db *sqlx.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (er *EquipmentRepository) Create(ctx context.Context, dto entities.EquipmentCreateDto) (int, error) {
	const op = "Repository/EquipmentRepository.Create"
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, status) values ($1, $2) RETURNING id", postgres.EquipmentTable)

	row := er.db.QueryRowContext(ctx, query, dto.Title, dto.Status)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: failed to execute equipment creation query: %w", op, err)
	}

	return id, nil
}

func (er *EquipmentRepository) GetAll(ctx context.Context) ([]entities.Equipment, error) {
	const op = "Repository/EquipmentRepository.GetAll"
	var equipment []entities.Equipment

	query := fmt.Sprintf(`
		SELECT e.id, e.title, e.status, e.created_at, 
       	CASE WHEN eu.equipment_id IS NULL AND r.equipment_id IS NULL THEN true ELSE false END AS is_available 
		FROM %s e LEFT JOIN %s eu ON e.id = eu.equipment_id AND current_date BETWEEN eu.start_date AND eu.end_date 
    	LEFT JOIN %s r ON e.id = r.equipment_id ORDER BY e.id ASC`,
		postgres.EquipmentTable, postgres.EquipmentUsageTable, postgres.ReportsTable)

	if err := er.db.SelectContext(ctx, &equipment, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting equipment: %w", op, ErrEquipmentNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting equipment: %w", op, err)
	}

	return equipment, nil
}

func (er *EquipmentRepository) GetAvailableByDate(ctx context.Context, dto entities.GetAvailableEquipmentByDateDto) ([]entities.Equipment, error) {
	const op = "Repository/EquipmentRepository.GetAvailableByDate"
	var equipment []entities.Equipment

	query := fmt.Sprintf(`
        SELECT e.* FROM %s e 
        WHERE e.id NOT IN (
            SELECT eu.equipment_id 
            FROM %s eu 
            WHERE 
                (eu.start_date <= $1 AND eu.end_date >= $1) OR 
                (eu.start_date <= $2 AND eu.end_date >= $2) OR 
                (eu.start_date >= $1 AND eu.end_date <= $2)
        ) 
        AND e.id NOT IN (
            SELECT r.equipment_id 
            FROM reports r
        ) 
        ORDER BY e.id ASC`,
		postgres.EquipmentTable, postgres.EquipmentUsageTable)

	if err := er.db.SelectContext(ctx, &equipment, query, dto.StartDate, dto.EndDate); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting equipment: %w", op, ErrEquipmentNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting equipment: %w", op, err)
	}

	return equipment, nil
}

func (er *EquipmentRepository) GetById(ctx context.Context, id int) (entities.Equipment, error) {
	const op = "Repository/EquipmentRepository.GetById"
	var equipment entities.Equipment

	query := fmt.Sprintf(`
		SELECT e.id, e.title, e.status, e.created_at,  
       	CASE 
       	    WHEN eu.equipment_id IS NULL 
       	             AND r.equipment_id IS NULL 
       	    THEN 
       	        true 
       	    ELSE 
       	        false 
       	    END AS is_available 
		FROM %s e LEFT JOIN %s eu ON e.id = eu.equipment_id AND CURRENT_DATE BETWEEN eu.start_date AND eu.end_date 
    	LEFT JOIN %s r ON e.id = r.equipment_id WHERE e.id = $1`,
		postgres.EquipmentTable, postgres.EquipmentUsageTable, postgres.ReportsTable)

	row := er.db.QueryRowContext(ctx, query, id)

	if err := row.Scan(&equipment.Id, &equipment.Title, &equipment.Status, &equipment.CreatedAt, &equipment.IsAvailable); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Equipment{}, fmt.Errorf("%s: error during getting equipment by id: %w", op, ErrEquipmentNotFound)
		}

		return entities.Equipment{}, fmt.Errorf("%s: error during getting equipment by id: %w", op, err)
	}

	return equipment, nil
}

func (er *EquipmentRepository) GetByEventId(ctx context.Context, eventID int) ([]entities.Equipment, error) {
	const op = "Repository/EquipmentRepository.GetByEventId"
	var equipment []entities.Equipment

	query := fmt.Sprintf(`
		SELECT e.* 
		FROM %s e
		JOIN %s eu ON eu.equipment_id = e.id
		WHERE eu.event_id = $1`,
		postgres.EquipmentTable, postgres.EquipmentUsageTable)
	if err := er.db.SelectContext(ctx, &equipment, query, eventID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting equipment by event id: %w", op, ErrEquipmentNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting equipment by event id: %w", op, ErrEquipmentNotFound)
	}

	return equipment, nil
}

func (er *EquipmentRepository) GetUsageHistoryById(ctx context.Context, id int) ([]entities.EquipmentUsageHistory, error) {
	const op = "Repository/EquipmentRepository.GetUsageHistoryById"
	dates := make([]entities.EquipmentUsageHistory, 0)

	query := fmt.Sprintf(`
		SELECT eu.id, u.username AS username, e.title AS event_title, eu.start_date, eu.end_date 
		FROM %s eu 
		JOIN users u ON u.id = eu.user_id 
    	JOIN events e ON e.id = eu.event_id 
        WHERE eu.equipment_id = $1 
        ORDER BY eu.start_date ASC`,
		postgres.EquipmentUsageTable)

	if err := er.db.SelectContext(ctx, &dates, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting equipment usage history by id: %w", op, ErrUsageHistoryNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting equipment usage history by id: %w", op, err)
	}

	return dates, nil
}

func (er *EquipmentRepository) GetEquipmentIsAvailableNow(ctx context.Context, id int) (bool, error) {
	const op = "Repository/EquipmentRepository.GetEquipmentIsAvailableNow"
	var count int
	query := fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM %s 
		WHERE equipment_id = $1 AND current_date BETWEEN start_date AND end_date`,
		postgres.EquipmentUsageTable)

	if err := er.db.QueryRowContext(ctx, query, id).Scan(&count); err != nil {
		return false, fmt.Errorf("%s: error during getting equipment usage history by id: %w", op, err)
	}

	return count == 0, nil
}

func (er *EquipmentRepository) Delete(ctx context.Context, id int) (int, error) {
	const op = "Repository/EquipmentRepository.Delete"
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.EquipmentTable)
	if _, err := er.db.ExecContext(ctx, query, id); err != nil {
		return 0, fmt.Errorf("%s: error during deleting equipment by id: %w", op, err)
	}

	return id, nil
}

func (er *EquipmentRepository) Update(ctx context.Context, id int, dto entities.EquipmentUpdateDto) error {
	const op = "Repository/EquipmentRepository.Update"
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *dto.Title)
		argId++
	}

	if dto.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *dto.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postgres.EquipmentTable, setQuery, argId)

	args = append(args, id)

	if _, err := er.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("%s: error during updating equipment by id: %w", op, ErrNotUpdated)
	}

	return nil
}
