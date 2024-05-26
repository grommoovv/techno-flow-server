package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (es *EventService) CreateEvent(dto domain.EventCreateDto) (int, error) {
	return es.repo.CreateEvent(dto)
}

func (es *EventService) GetAllEvents() ([]domain.Event, error) {
	return es.repo.GetAllEvents()
}

func (es *EventService) GetEventById(id int) (domain.Event, error) {
	return es.repo.GetEventById(id)
}

func (es *EventService) DeleteEvent(id int) (int, error) {
	return es.repo.DeleteEvent(id)
}

func (es *EventService) UpdateEvent() {}
