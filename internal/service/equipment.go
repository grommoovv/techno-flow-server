package service

import (
	"server-techno-flow/internal/repository"
)

type EquipmentService struct {
	repo repository.Equipment
}

func NewEquipmentService(repo repository.Equipment) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (es *EquipmentService) CreateEquipment() {}

func (es *EquipmentService) GetEquipment() {}

func (es *EquipmentService) GetAllEquipment() {}

func (es *EquipmentService) DeleteEquipment() {}

func (es *EquipmentService) UpdateEquipment() {}
