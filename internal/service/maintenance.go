package service

import (
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

func (ms *MaintenanceService) Create(dto entities.MaintenanceCreateDto) (int, error) {
	fixedIn := utils.GenerageRandomDate(ms.random)
	dto.FixedIn = fixedIn
	return ms.repo.Create(dto)
}

func (ms *MaintenanceService) GetAll() ([]entities.Maintenance, error) {
	return ms.repo.GetAll()
}

func (ms *MaintenanceService) GetById(id int) (entities.Maintenance, error) {
	return ms.repo.GetById(id)
}

func (ms *MaintenanceService) Delete(id int) error {
	return ms.repo.Delete(id)
}

func (ms *MaintenanceService) Update() {}
