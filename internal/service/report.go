package service

import (
	"github.com/sirupsen/logrus"
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type ReportService struct {
	repo               repository.Report
	equipmentService   EquipmentService
	maintenanceService MaintenanceService
}

func NewReportService(repo repository.Report, EquipmentService *EquipmentService, MaintenanceService *MaintenanceService) *ReportService {
	return &ReportService{
		repo:               repo,
		equipmentService:   *EquipmentService,
		maintenanceService: *MaintenanceService,
	}
}

func (rs *ReportService) CreateReport(dto entities.ReportCreateDto) (int, error) {
	var updateEquipmentDto entities.EquipmentUpdateDto
	status := "На обслуживании"
	updateEquipmentDto.Status = &status

	equipmentID := dto.EquipmentId

	if err := rs.equipmentService.UpdateEquipment(equipmentID, updateEquipmentDto); err != nil {
		logrus.Errorf("error updating equipment status: %v", err.Error())
		return 0, err
	}

	var createMaintenanceDto entities.MaintenanceCreateDto
	createMaintenanceDto.EquipmentId = equipmentID

	_, err := rs.maintenanceService.Create(createMaintenanceDto)

	if err != nil {
		logrus.Errorf("error creating maintenance: %v", err.Error())
		return 0, err
	}

	reportID, err := rs.repo.Create(dto)

	if err != nil {
		logrus.Errorf("error creating report: %v", err.Error())
		return 0, err
	}

	return reportID, nil
}

func (rs *ReportService) GetAllReports() ([]entities.Report, error) {
	return rs.repo.GetAll()
}

func (rs *ReportService) GetReportById(id int) (entities.Report, error) {
	return rs.repo.GetById(id)
}

func (rs *ReportService) GetReportsByUserId(id int) ([]entities.Report, error) {
	return rs.repo.GetByUserId(id)
}

func (rs *ReportService) DeleteReport(id int) error {
	return rs.repo.Delete(id)
}

func (rs *ReportService) UpdateReport() {}
