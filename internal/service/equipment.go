package service

import (
	"context"
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type EquipmentService struct {
	repo repository.Equipment
}

func NewEquipmentService(repo repository.Equipment) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (es *EquipmentService) CreateEquipment(ctx context.Context, dto entities.EquipmentCreateDto) (int, error) {
	return es.repo.Create(ctx, dto)
}

func (es *EquipmentService) GetAllEquipment(ctx context.Context) ([]entities.Equipment, error) {
	return es.repo.GetAll(ctx)
}

func (es *EquipmentService) GetAvailableEquipmentByDate(ctx context.Context, dto entities.GetAvailableEquipmentByDateDto) ([]entities.Equipment, error) {
	return es.repo.GetAvailableByDate(ctx, dto)
}

func (es *EquipmentService) GetEquipmentById(ctx context.Context, id int) (entities.Equipment, error) {
	return es.repo.GetById(ctx, id)
}

func (es *EquipmentService) GetEquipmentByEventId(ctx context.Context, id int) ([]entities.Equipment, error) {
	return es.repo.GetByEventId(ctx, id)
}

func (es *EquipmentService) GetEquipmentUsageHistoryById(ctx context.Context, id int) ([]entities.EquipmentUsageHistory, error) {
	return es.repo.GetUsageHistoryById(ctx, id)
}

func (es *EquipmentService) DeleteEquipment(ctx context.Context, id int) (int, error) {
	return es.repo.Delete(ctx, id)
}

func (es *EquipmentService) UpdateEquipment(ctx context.Context, id int, dto entities.EquipmentUpdateDto) error {
	return es.repo.Update(ctx, id, dto)
}
