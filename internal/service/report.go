package service

import (
	"github.com/sirupsen/logrus"
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type ReportService struct {
	repo             repository.Report
	equipmentService EquipmentService
}

func NewReportService(repo repository.Report, EquipmentService *EquipmentService) *ReportService {
	return &ReportService{repo: repo, equipmentService: *EquipmentService}
}

func (rs *ReportService) CreateReport(dto domain.ReportCreateDto) (int, error) {
	var updateEquipmentDto domain.EquipmentUpdateDto
	status := "На обслуживании"
	updateEquipmentDto.Status = &status

	equipmentID := dto.EquipmentId

	if err := rs.equipmentService.UpdateEquipment(equipmentID, updateEquipmentDto); err != nil {
		logrus.Errorf("error updating equipment status: %v", err.Error())
		return 0, err
	}

	reportID, err := rs.repo.CreateReport(dto)

	return reportID, err
}

func (rs *ReportService) GetAllReports() ([]domain.Report, error) {
	return rs.repo.GetAllReports()
}

func (rs *ReportService) GetReportById(id int) (domain.Report, error) {
	return rs.repo.GetReportById(id)
}

func (rs *ReportService) DeleteReport(id int) (int, error) {
	return rs.repo.DeleteReport(id)
}

func (rs *ReportService) UpdateReport() {}
