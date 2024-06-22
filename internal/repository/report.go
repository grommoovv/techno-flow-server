package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/entities"
)

var (
	ErrReportNotFound  = errors.New("report not found")
	ErrReportsNotFound = errors.New("reports not found")
	ErrNotCreated      = errors.New("failed to create report")
)

type ReportRepository struct {
	db *sqlx.DB
}

func NewReportRepository(db *sqlx.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (rr *ReportRepository) Create(ctx context.Context, dto entities.ReportCreateDto) (int, error) {
	const op = "Repository/ReportRepository.Create"
	var id int

	query := fmt.Sprintf("INSERT INTO %s (message, user_id, equipment_id) values ($1, $2, $3) RETURNING id", postgres.ReportsTable)

	row := rr.db.QueryRowContext(ctx, query, dto.Message, dto.UserId, dto.EquipmentId)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: error during creating report: %w", op, ErrNotCreated)
	}

	return id, nil
}

func (rr *ReportRepository) GetAll(ctx context.Context) ([]entities.Report, error) {
	const op = "Repository/ReportRepository.GetAll"
	var reports []entities.Report

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", postgres.ReportsTable)
	if err := rr.db.SelectContext(ctx, &reports, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: error during getting reports: %w", op, ErrReportsNotFound)
		}

		return nil, fmt.Errorf("%s: error during getting reports: %w", op, err)
	}

	return reports, nil
}

func (rr *ReportRepository) GetById(ctx context.Context, id int) (entities.Report, error) {
	const op = "Repository/ReportRepository.GetByUserId"
	var report entities.Report

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.ReportsTable)

	row := rr.db.QueryRowContext(ctx, query, id)

	if err := row.Scan(&report.Id, &report.Message, &report.CreatedAt, &report.UserId, &report.EquipmentId); err != nil {
		return report, fmt.Errorf("%s: error during getting report by id: %w", op, ErrReportNotFound)
	}

	return report, nil
}

func (rr *ReportRepository) GetByUserId(ctx context.Context, id int) ([]entities.Report, error) {
	const op = "Repository/ReportRepository.GetByUserId"
	var reports []entities.Report

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", postgres.ReportsTable)

	if err := rr.db.SelectContext(ctx, &reports, query, id); err != nil {
		return nil, fmt.Errorf("%s: error during getting report by user id: %w", op, err)
	}

	return reports, nil
}

func (rr *ReportRepository) Delete(ctx context.Context, id int) error {
	const op = "Repository/ReportRepository.Create"

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.ReportsTable)

	if _, err := rr.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("%s: error during deleting report by id: %w", op, err)
	}

	return nil
}

func (rr *ReportRepository) Update(ctx context.Context) {}
