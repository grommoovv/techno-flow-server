package service

import (
	"server-techno-flow/internal/repository"
)

type ReportService struct {
	repo repository.Report
}

func NewReportService(repo repository.Report) *ReportService {
	return &ReportService{repo: repo}
}

func (rs *ReportService) CreateReport() {}

func (rs *ReportService) GetReport() {}

func (rs *ReportService) GetAllReports() {}

func (rs *ReportService) DeleteReport() {}

func (rs *ReportService) UpdateReport() {}
