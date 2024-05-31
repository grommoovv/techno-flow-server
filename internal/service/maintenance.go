package service

import (
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type MaintenanceService struct {
	repo repository.Maintenance
}

func NewMaintenanceService(repo repository.Maintenance) *MaintenanceService {
	return &MaintenanceService{repo: repo}
}

func (ms *MaintenanceService) Create(dto entities.MaintenanceCreateDto) (int, error) {
	return ms.Create(dto)
}

func (ms *MaintenanceService) GetAll() ([]entities.Maintenance, error) {
	return ms.repo.GetAll()
}

func (ms *MaintenanceService) GetById(id int) (entities.Maintenance, error) {
	return ms.repo.GetById(id)
}

func (ms *MaintenanceService) Delete() {}

func (ms *MaintenanceService) Update() {}
