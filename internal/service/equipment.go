package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type EquipmentService struct {
	repo repository.Equipment
}

func NewEquipmentService(repo repository.Equipment) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (es *EquipmentService) CreateEquipment(dto domain.EquipmentCreateDto) (int, error) {
	return es.repo.Create(dto)
}

func (es *EquipmentService) GetAllEquipment() ([]domain.Equipment, error) {
	return es.repo.GetAll()
}

func (es *EquipmentService) GetAvailableEquipment() ([]domain.Equipment, error) {
	return es.repo.GetAvailable()
}

func (es *EquipmentService) GetEquipmentById(id int) (domain.Equipment, error) {
	return es.repo.GetById(id)
}

func (es *EquipmentService) GetEquipmentUsageHistoryById(id int) ([]domain.EquipmentUsageHistory, error) {
	return es.repo.GetUsageHistoryById(id)
}

func (es *EquipmentService) DeleteEquipment(id int) (int, error) {
	return es.repo.Delete(id)
}

func (es *EquipmentService) UpdateEquipment(id int, dto domain.EquipmentUpdateDto) error {
	return es.repo.Update(id, dto)
}
