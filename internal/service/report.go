package service

import (
	"github.com/sirupsen/logrus"
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type ReportService struct {
	repo             repository.Report
	equipmentService EquipmentService
}

func NewReportService(repo repository.Report, EquipmentService *EquipmentService) *ReportService {
	return &ReportService{repo: repo, equipmentService: *EquipmentService}
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

	reportID, err := rs.repo.CreateReport(dto)

	return reportID, err
}

func (rs *ReportService) GetAllReports() ([]entities.Report, error) {
	return rs.repo.GetAllReports()
}

func (rs *ReportService) GetReportById(id int) (entities.Report, error) {
	return rs.repo.GetReportById(id)
}

func (rs *ReportService) DeleteReport(id int) error {
	return rs.repo.DeleteReport(id)
}

func (rs *ReportService) UpdateReport() {}
