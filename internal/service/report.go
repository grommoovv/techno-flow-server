package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type ReportService struct {
	repo repository.Report
}

func NewReportService(repo repository.Report) *ReportService {
	return &ReportService{repo: repo}
}

func (rs *ReportService) CreateReport(dto domain.ReportCreateDto) (int, error) {
	return rs.repo.CreateReport(dto)
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
