package service

import (
	"server-techno-flow/internal/repository"
)

type MaintenanceService struct {
	repo repository.Maintenance
}

func NewMaintenanceService(repo repository.Maintenance) *MaintenanceService {
	return &MaintenanceService{repo: repo}
}

func (ms *MaintenanceService) CreateMaintenance() {}

func (ms *MaintenanceService) GetMaintenance() {}

func (ms *MaintenanceService) GetAllMaintenance() {}

func (ms *MaintenanceService) DeleteMaintenance() {}

func (ms *MaintenanceService) UpdateMaintenance() {}
