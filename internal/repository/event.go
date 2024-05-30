package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/domain"
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) Create(dto domain.EventCreateDto) (int, error) {
	var eventID int

	for _, equipmentID := range dto.EquipmentIDs {
		var count int

		query := fmt.Sprintf("SELECT COUNT(*) FROM equipment_usage WHERE equipment_id = $1 AND (start_date <= $2 AND end_date >= $2 OR start_date <= $3 AND end_date >= $3 OR start_date >= $2 AND end_date <= $3)")

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

func (er *EventRepository) GetAll() ([]domain.Event, error) {
	var events []domain.Event
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.EventsTable)
	if err := er.db.Select(&events, query); err != nil {
		return nil, err
	}
	return events, nil
}

func (er *EventRepository) GetById(id int) (domain.Event, error) {
	var event domain.Event

	query := fmt.Sprintf("SELECT id, title, type, start_date, end_date, duration, status, user_id FROM %s WHERE id = $1", postgres.EventsTable)

	err := er.db.QueryRow(query, id).Scan(&event.ID, &event.Title, &event.Type, &event.StartDate, &event.EndDate, &event.Duration, &event.Status, &event.UserID)

	return event, err
}

func (er *EventRepository) Delete(id int) (int, error) {

	query := fmt.Sprintf("SELECT equipment_id FROM %s WHERE event_id=$1", postgres.EquipmentUsageTable)
	rows, err := er.db.Query(query, id)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var equipmentID int
		err := rows.Scan(&equipmentID)
		if err != nil {
			return 0, err
		}

		query = fmt.Sprintf("UPDATE %s SET reserved_at=NULL WHERE id=$1", postgres.EquipmentTable)
		if _, err := er.db.Exec(query, equipmentID); err != nil {
			return 0, err
		}
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE event_id=$1", postgres.EquipmentUsageTable)
	if _, err := er.db.Exec(query, id); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.EventsTable)
	if _, err := er.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *EventRepository) Update() {}
