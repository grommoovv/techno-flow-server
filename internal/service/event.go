package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type EventService struct {
	repo repository.Event
	EquipmentService
}

func NewEventService(repo repository.Event, equipmentService *EquipmentService) *EventService {
	return &EventService{repo: repo, EquipmentService: *equipmentService}
}

func (es *EventService) CreateEvent(dto domain.EventCreateDto) (int, error) {
	//var updateEquipmentDto domain.EquipmentUpdateDto
	//updateEquipmentDto.ReservedAt = &dto.StartDate
	//
	//for _, equipID := range dto.EquipmentId {
	//	if err := es.UpdateEquipment(equipID, updateEquipmentDto); err != nil {
	//		return 0, err
	//	}
	//}

	eventID, err := es.repo.Create(dto)

	return eventID, err
}

func (es *EventService) GetAllEvents() ([]domain.Event, error) {
	return es.repo.GetAll()
}

func (es *EventService) GetEventById(id int) (domain.Event, error) {
	return es.repo.GetById(id)
}

func (es *EventService) GetEventsByUserId(id int) ([]domain.Event, error) {
	return es.repo.GetByUserId(id)
}

func (es *EventService) DeleteEvent(id int) (int, error) {
	//var updateEquipmentDto domain.EquipmentUpdateDto
	//updateEquipmentDto.ReservedAt = &dto.StartDate
	//
	//for _, equipID := range dto.EquipmentId {
	//	if err := es.UpdateEquipment(equipID, updateEquipmentDto); err != nil {
	//		return 0, err
	//	}
	//}

	eventID, err := es.repo.Delete(id)

	return eventID, err
}

func (es *EventService) UpdateEvent() {}
