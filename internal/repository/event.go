package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"

	"github.com/jmoiron/sqlx"
)

var (
	ErrCreateEvent            = errors.New("failed to create event")
	ErrEventsNotFound         = errors.New("events not found")
	ErrEventNotFound          = errors.New("event not found")
	ErrEquipmentUsageNotFound = errors.New("equipment usage not found")
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) Create(ctx context.Context, dto entities.EventCreateDto) (int, error) {

	const op = "Repository/EventRepository.Create"

	var eventID int

	for _, equipmentID := range dto.EquipmentIDs {
		var count int

		q := fmt.Sprintf(`
			SELECT COUNT(*) 
			FROM %s eu 
			LEFT JOIN %s r ON eu.equipment_id = r.equipment_id 
			WHERE eu.equipment_id = $1 
			AND r.equipment_id IS NULL AND ( eu.start_date <= $2 AND eu.end_date >= $2 
			OR eu.start_date <= $3 AND eu.end_date >= $3 OR eu.start_date >= $2 AND eu.end_date <= $3)`,
			postgres.EquipmentUsageTable, postgres.ReportsTable)

		if err := er.db.QueryRowContext(ctx, q, equipmentID, dto.StartDate, dto.EndDate).Scan(&count); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return 0, fmt.Errorf("%s: error during getting equipment usage: %w", op, ErrEquipmentUsageNotFound)
			}
			return 0, fmt.Errorf("%s: error during getting equipment usage: %w", op, err)
		}

		if count != 0 {
			return 0, fmt.Errorf("equipment with ID: %d is already in use at this time", equipmentID)
		}
	}

	q := fmt.Sprintf("INSERT INTO %s (title, type, start_date, end_date, user_id) values ($1, $2, $3, $4, $5) RETURNING id", postgres.EventsTable)
	row := er.db.QueryRowContext(ctx, q, dto.Title, dto.Type, dto.StartDate, dto.EndDate, dto.UserID)
	if err := row.Scan(&eventID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("%s: error during creating event: %w", op, ErrCreateEvent)
		}
		return 0, fmt.Errorf("%s: error during getting event: %w", op, err)
	}

	for _, equipmentID := range dto.EquipmentIDs {
		q = fmt.Sprintf("INSERT INTO %s (user_id, event_id, equipment_id, start_date, end_date) values ($1, $2, $3, $4, $5)", postgres.EquipmentUsageTable)
		_, err := er.db.ExecContext(ctx, q, dto.UserID, eventID, equipmentID, dto.StartDate, dto.EndDate)
		if err != nil {
			return 0, fmt.Errorf("%s: failed to execute equipment usage creation query: %w", op, err)
		}
	}

	return eventID, nil
}

func (er *EventRepository) GetAll(ctx context.Context) ([]entities.Event, error) {
	const op = "Repository/EventRepository.GetAll"

	var events []entities.Event

	q := fmt.Sprintf(`
		SELECT e.id, e.title, e.type, e.start_date, e.end_date,  e.duration, e.status, e.user_id, u.username 
		FROM %s e 
		JOIN %s u on e.user_id = u.id 
        ORDER BY e.start_date ASC`,
		postgres.EventsTable, postgres.UsersTable)

	if err := er.db.SelectContext(ctx, &events, q); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting events: %w", op, ErrEventsNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting events: %w", op, err)
	}

	return events, nil
}

func (er *EventRepository) GetById(ctx context.Context, id int) (entities.Event, error) {
	const op = "Repository/EventRepository.GetById"

	var event entities.Event

	q := fmt.Sprintf(`
        SELECT e.id, e.title, e.type, e.start_date, e.end_date, e.duration, e.status, e.user_id, u.username
        FROM %s e
        JOIN %s u ON e.user_id = u.id
        WHERE e.id = $1`,
		postgres.EventsTable, postgres.UsersTable)

	row := er.db.QueryRowContext(ctx, q, id)

	if err := row.Scan(&event.ID, &event.Title, &event.Type, &event.StartDate, &event.EndDate, &event.Duration, &event.Status, &event.UserID, &event.Username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Event{}, fmt.Errorf("%s: error during getting event: %w", op, ErrEventNotFound)
		}

		return entities.Event{}, fmt.Errorf("%s: error during getting event: %w", op, err)
	}

	return event, nil
}

func (er *EventRepository) GetByUserId(ctx context.Context, id int) ([]entities.Event, error) {
	const op = "Repository/EventRepository.GetByUserId"

	var events []entities.Event

	q := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", postgres.EventsTable)

	if err := er.db.SelectContext(ctx, &events, q, id); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting user's events: %w", op, ErrEventsNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting user's events: %w", op, err)
	}

	return events, nil
}

func (er *EventRepository) Delete(ctx context.Context, id int) error {
	const op = "Repository/EventRepository.Delete"

	q := fmt.Sprintf("SELECT equipment_id FROM %s WHERE event_id=$1", postgres.EquipmentUsageTable)
	rows, err := er.db.QueryContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: error during getting equipment usage: %w", op, ErrEquipmentUsageNotFound)
	}

	defer rows.Close()

	for rows.Next() {
		var equipmentID int
		if err := rows.Scan(&equipmentID); err != nil {
			return err
		}

		q = fmt.Sprintf("UPDATE %s SET reserved_at=NULL WHERE id=$1", postgres.EquipmentTable)
		if _, err := er.db.ExecContext(ctx, q, equipmentID); err != nil {
			return err
		}
	}

	q = fmt.Sprintf("DELETE FROM %s WHERE event_id=$1", postgres.EquipmentUsageTable)
	if _, err := er.db.ExecContext(ctx, q, id); err != nil {
		return err
	}

	q = fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.EventsTable)
	if _, err := er.db.ExecContext(ctx, q, id); err != nil {
		return err
	}

	return nil
}

func (er *EventRepository) Update(ctx context.Context) {}
