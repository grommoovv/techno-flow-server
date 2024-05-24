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
	return es.repo.CreateEquipment(dto)
}

func (es *EquipmentService) GetAllEquipment() ([]domain.Equipment, error) {
	return es.repo.GetAllEquipment()
}

func (es *EquipmentService) GetEquipmentById(id int) (domain.Equipment, error) {
	return es.repo.GetEquipmentById(id)
}

func (es *EquipmentService) DeleteEquipment(id int) (int, error) {
	return es.repo.DeleteEquipment(id)
}

func (es *EquipmentService) UpdateEquipment(id int, dto domain.EquipmentUpdateDto) error {
	return es.repo.UpdateEquipment(id, dto)
}
