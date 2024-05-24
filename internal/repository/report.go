package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/domain"
)

type ReportRepository struct {
	db *sqlx.DB
}

func NewReportRepository(db *sqlx.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (rr *ReportRepository) CreateReport(dto domain.ReportCreateDto) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (message, user_id, equipment_id) values ($1, $2, $3) RETURNING id", postgres.ReportsTable)

	row := rr.db.QueryRow(query, dto.Message, dto.UserId, dto.EquipmentId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (rr *ReportRepository) GetAllReports() ([]domain.Report, error) {
	var reports []domain.Report
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.ReportsTable)
	if err := rr.db.Select(&reports, query); err != nil {
		return nil, err
	}
	return reports, nil
}

func (rr *ReportRepository) GetReportById(id int) (domain.Report, error) {
	var report domain.Report

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.ReportsTable)

	err := rr.db.QueryRow(query, id).Scan(&report.Id, &report.Message, &report.CreatedAt, &report.UserId, &report.EquipmentId)

	return report, err
}

func (rr *ReportRepository) DeleteReport(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.ReportsTable)
	if _, err := rr.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (rr *ReportRepository) UpdateReport() {}
