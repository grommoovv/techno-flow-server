package service

import (
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type EquipmentService struct {
	repo repository.Equipment
}

func NewEquipmentService(repo repository.Equipment) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (es *EquipmentService) CreateEquipment(dto entities.EquipmentCreateDto) (int, error) {
	return es.repo.Create(dto)
}

func (es *EquipmentService) GetAllEquipment() ([]entities.Equipment, error) {
	return es.repo.GetAll()
}

func (es *EquipmentService) GetAvailableEquipmentByDate(dto entities.GetAvailableEquipmentByDateDto) ([]entities.Equipment, error) {
	return es.repo.GetAvailableByDate(dto)
}

func (es *EquipmentService) GetEquipmentById(id int) (entities.Equipment, error) {
	return es.repo.GetById(id)
}

func (es *EquipmentService) GetEquipmentUsageHistoryById(id int) ([]entities.EquipmentUsageHistory, error) {
	return es.repo.GetUsageHistoryById(id)
}

func (es *EquipmentService) DeleteEquipment(id int) (int, error) {
	return es.repo.Delete(id)
}

func (es *EquipmentService) UpdateEquipment(id int, dto entities.EquipmentUpdateDto) error {
	return es.repo.Update(id, dto)
}
