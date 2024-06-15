package service

import (
	"context"
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

func (rs *ReportService) CreateReport(ctx context.Context, dto entities.ReportCreateDto) (int, error) {
	var updateEquipmentDto entities.EquipmentUpdateDto
	status := "На обслуживании"
	updateEquipmentDto.Status = &status

	equipmentID := dto.EquipmentId

	if err := rs.equipmentService.UpdateEquipment(ctx, equipmentID, updateEquipmentDto); err != nil {
		logrus.Errorf("error updating equipment status: %v", err.Error())
		return 0, err
	}

	var createMaintenanceDto entities.MaintenanceCreateDto
	createMaintenanceDto.EquipmentId = equipmentID

	_, err := rs.maintenanceService.Create(ctx, createMaintenanceDto)

	if err != nil {
		logrus.Errorf("error creating maintenance: %v", err.Error())
		return 0, err
	}

	reportID, err := rs.repo.Create(ctx, dto)

	if err != nil {
		logrus.Errorf("error creating report: %v", err.Error())
		return 0, err
	}

	return reportID, nil
}

func (rs *ReportService) GetAllReports(ctx context.Context) ([]entities.Report, error) {
	return rs.repo.GetAll(ctx)
}

func (rs *ReportService) GetReportById(ctx context.Context, id int) (entities.Report, error) {
	return rs.repo.GetById(ctx, id)
}

func (rs *ReportService) GetReportsByUserId(ctx context.Context, id int) ([]entities.Report, error) {
	return rs.repo.GetByUserId(ctx, id)
}

func (rs *ReportService) DeleteReport(ctx context.Context, id int) error {
	return rs.repo.Delete(ctx, id)
}

func (rs *ReportService) UpdateReport(ctx context.Context) {}
