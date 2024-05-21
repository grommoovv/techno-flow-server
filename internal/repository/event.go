package repository

import (
	"github.com/jmoiron/sqlx"
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) CreateEvent() {}

func (er *EventRepository) GetEvent() {}

func (er *EventRepository) GetAllEvents() {}

func (er *EventRepository) DeleteEvent() {}

func (er *EventRepository) UpdateEvent() {}
