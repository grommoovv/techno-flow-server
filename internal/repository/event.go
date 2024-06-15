package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) Create(ctx context.Context, dto entities.EventCreateDto) (int, error) {
	var eventID int

	for _, equipmentID := range dto.EquipmentIDs {
		var count int

		query := fmt.Sprintf("SELECT COUNT(*) FROM %s eu LEFT JOIN %s r ON eu.equipment_id = r.equipment_id WHERE eu.equipment_id = $1 AND r.equipment_id IS NULL AND ( eu.start_date <= $2 AND eu.end_date >= $2 OR eu.start_date <= $3 AND eu.end_date >= $3 OR eu.start_date >= $2 AND eu.end_date <= $3)", postgres.EquipmentUsageTable, postgres.ReportsTable)

		if err := er.db.QueryRow(query, equipmentID, dto.StartDate, dto.EndDate).Scan(&count); err != nil {
			return 0, err
		}

		if count != 0 {
			return 0, fmt.Errorf("equipment with ID %d is already in use at this time", equipmentID)
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (title, type, start_date, end_date, user_id) values ($1, $2, $3, $4, $5) RETURNING id", postgres.EventsTable)
	row := er.db.QueryRow(query, dto.Title, dto.Type, dto.StartDate, dto.EndDate, dto.UserID)
	if err := row.Scan(&eventID); err != nil {
		return 0, err
	}

	fmt.Printf("equipmentId value: %v, type: %T\n", dto.EquipmentIDs, dto.EquipmentIDs)

	for _, equipmentID := range dto.EquipmentIDs {
		query := fmt.Sprintf("INSERT INTO %s (user_id, event_id, equipment_id, start_date, end_date) values ($1, $2, $3, $4, $5)", postgres.EquipmentUsageTable)
		_, err := er.db.Exec(query, dto.UserID, eventID, equipmentID, dto.StartDate, dto.EndDate)
		if err != nil {
			return 0, err
		}
	}

	return eventID, nil
}

func (er *EventRepository) GetAll(ctx context.Context) ([]entities.Event, error) {
	var events []entities.Event
	query := fmt.Sprintf("SELECT e.id, e.title, e.type, e.start_date, e.end_date,  e.duration, e.status, e.user_id, u.username FROM %s e JOIN %s u on e.user_id = u.id ORDER BY e.start_date ASC", postgres.EventsTable, postgres.UsersTable)
	if err := er.db.Select(&events, query); err != nil {
		return nil, err
	}
	return events, nil
}

func (er *EventRepository) GetById(ctx context.Context, id int) (entities.Event, error) {
	var event entities.Event

	query := fmt.Sprintf(`
        SELECT e.id, e.title, e.type, e.start_date, e.end_date, e.duration, e.status, e.user_id, u.username
        FROM %s e
        JOIN %s u ON e.user_id = u.id
        WHERE e.id = $1`,
		postgres.EventsTable, postgres.UsersTable)

	err := er.db.QueryRow(query, id).Scan(&event.ID, &event.Title, &event.Type, &event.StartDate, &event.EndDate, &event.Duration, &event.Status, &event.UserID, &event.Username)

	return event, err
}

func (er *EventRepository) GetByUserId(ctx context.Context, id int) ([]entities.Event, error) {
	var events []entities.Event

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", postgres.EventsTable)

	if err := er.db.Select(&events, query, id); err != nil {
		return nil, err
	}

	return events, nil
}

func (er *EventRepository) Delete(ctx context.Context, id int) error {

	query := fmt.Sprintf("SELECT equipment_id FROM %s WHERE event_id=$1", postgres.EquipmentUsageTable)
	rows, err := er.db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var equipmentID int
		err := rows.Scan(&equipmentID)
		if err != nil {
			return err
		}

		query = fmt.Sprintf("UPDATE %s SET reserved_at=NULL WHERE id=$1", postgres.EquipmentTable)
		if _, err := er.db.Exec(query, equipmentID); err != nil {
			return err
		}
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE event_id=$1", postgres.EquipmentUsageTable)
	if _, err := er.db.Exec(query, id); err != nil {
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.EventsTable)
	if _, err := er.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (er *EventRepository) Update(ctx context.Context) {}
