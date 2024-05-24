package service

import (
	"server-techno-flow/internal/repository"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (es *EventService) CreateEvent() {}

func (es *EventService) GetEventById() {}

func (es *EventService) GetAllEvents() {}

func (es *EventService) DeleteEvent() {}

func (es *EventService) UpdateEvent() {}
