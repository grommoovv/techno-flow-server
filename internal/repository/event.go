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

func (er *EventRepository) CreateEvent(dto domain.EventCreateDto) (int, error) {
	var eventId int

	query := fmt.Sprintf("INSERT INTO %s (title, type, start_date, end_date, user_id) values ($1, $2, $3, $4, $5) RETURNING id", postgres.EventsTable)
	row := er.db.QueryRow(query, dto.Title, dto.Type, dto.StartDate, dto.EndDate, dto.UserId)
	if err := row.Scan(&eventId); err != nil {
		return 0, err
	}

	fmt.Printf("EquipmentId value: %v, type: %T\n", dto.EquipmentId, dto.EquipmentId)

	if equipmentID, ok := dto.EquipmentId.(float64); ok {
		query := fmt.Sprintf("INSERT INTO %s (user_id, event_id, equipment_id) values ($1, $2, $3)", postgres.EquipmentUsageTable)
		_, err := er.db.Exec(query, dto.UserId, eventId, equipmentID)
		if err != nil {
			return 0, err
		}
	} else if equipmentIDs, ok := dto.EquipmentId.([]interface{}); ok {
		for _, eqID := range equipmentIDs {
			query := fmt.Sprintf("INSERT INTO %s (user_id, event_id, equipment_id) values ($1, $2, $3)", postgres.EquipmentUsageTable)
			_, err := er.db.Exec(query, dto.UserId, eventId, eqID)
			if err != nil {
				return 0, err
			}
		}
	}

	return eventId, nil
}

func (er *EventRepository) GetAllEvents() ([]domain.Event, error) {
	var events []domain.Event
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.EventsTable)
	if err := er.db.Select(&events, query); err != nil {
		return nil, err
	}
	return events, nil
}

func (er *EventRepository) GetEventById(id int) (domain.Event, error) {
	var event domain.Event

	query := fmt.Sprintf("SELECT id, title, type, start_date, end_date, duration, status, user_id FROM %s WHERE id = $1", postgres.EventsTable)

	err := er.db.QueryRow(query, id).Scan(&event.Id, &event.Title, &event.Type, &event.StartDate, &event.EndDate, &event.Duration, &event.Status, &event.UserId)

	return event, err
}

func (er *EventRepository) DeleteEvent(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.EventsTable)
	if _, err := er.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *EventRepository) UpdateEvent() {}
