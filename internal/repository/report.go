package repository

import (
	"github.com/jmoiron/sqlx"
)

type ReportRepository struct {
	db *sqlx.DB
}

func NewReportRepository(db *sqlx.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (rr *ReportRepository) CreateReport() {}

func (rr *ReportRepository) GetReport() {}

func (rr *ReportRepository) GetAllReports() {}

func (rr *ReportRepository) DeleteReport() {}

func (rr *ReportRepository) UpdateReport() {}
