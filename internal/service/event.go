package service

import (
	"context"
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type EventService struct {
	repo repository.Event
	EquipmentService
}

func NewEventService(repo repository.Event, equipmentService *EquipmentService) *EventService {
	return &EventService{repo: repo, EquipmentService: *equipmentService}
}

func (es *EventService) CreateEvent(ctx context.Context, dto entities.EventCreateDto) (int, error) {
	//var updateEquipmentDto entities.EquipmentUpdateDto
	//updateEquipmentDto.ReservedAt = &dto.StartDate
	//
	//for _, equipID := range dto.EquipmentId {
	//	if err := es.UpdateEquipment(equipID, updateEquipmentDto); err != nil {
	//		return 0, err
	//	}
	//}

	eventID, err := es.repo.Create(ctx, dto)

	return eventID, err
}

func (es *EventService) GetAllEvents(ctx context.Context) ([]entities.Event, error) {
	return es.repo.GetAll(ctx)
}

func (es *EventService) GetEventById(ctx context.Context, id int) (entities.Event, error) {
	return es.repo.GetById(ctx, id)
}

func (es *EventService) GetEventsByUserId(ctx context.Context, id int) ([]entities.Event, error) {
	return es.repo.GetByUserId(ctx, id)
}

func (es *EventService) DeleteEvent(ctx context.Context, id int) error {
	//var updateEquipmentDto entities.EquipmentUpdateDto
	//updateEquipmentDto.ReservedAt = &dto.StartDate
	//
	//for _, equipID := range dto.EquipmentId {
	//	if err := es.UpdateEquipment(equipID, updateEquipmentDto); err != nil {
	//		return 0, err
	//	}
	//}

	err := es.repo.Delete(ctx, id)

	return err
}

func (es *EventService) UpdateEvent(ctx context.Context) {}
