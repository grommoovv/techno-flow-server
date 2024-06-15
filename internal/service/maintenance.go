package service

import (
	"context"
	"math/rand"
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
	"server-techno-flow/internal/utils"
)

type MaintenanceService struct {
	repo   repository.Maintenance
	random *rand.Rand
}

func NewMaintenanceService(repo repository.Maintenance, random *rand.Rand) *MaintenanceService {
	return &MaintenanceService{repo: repo, random: random}
}

func (ms *MaintenanceService) Create(ctx context.Context, dto entities.MaintenanceCreateDto) (int, error) {
	fixedIn := utils.GenerageRandomDate(ms.random)
	dto.FixedIn = fixedIn
	return ms.repo.Create(ctx, dto)
}

func (ms *MaintenanceService) GetAll(ctx context.Context) ([]entities.Maintenance, error) {
	return ms.repo.GetAll(ctx)
}

func (ms *MaintenanceService) GetById(ctx context.Context, id int) (entities.Maintenance, error) {
	return ms.repo.GetById(ctx, id)
}

func (ms *MaintenanceService) Delete(ctx context.Context, id int) error {
	return ms.repo.Delete(ctx, id)
}

func (ms *MaintenanceService) Update(ctx context.Context) {}
