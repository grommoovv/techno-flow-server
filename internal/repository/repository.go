package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/entities"
)

type (
	User interface {
		GetAll(ctx context.Context) ([]entities.User, error)
		GetById(ctx context.Context, id int) (entities.User, error)
		GetByCredentials(ctx context.Context, dto entities.UserSignInDto) (entities.User, error)

		Create(ctx context.Context, dto entities.UserCreateDto) (int, error)
		Delete(ctx context.Context, id int) (int, error)
		Update(ctx context.Context, id int, dto entities.UserUpdateDto) error
	}

	Equipment interface {
		GetAll(ctx context.Context) ([]entities.Equipment, error)
		GetAvailableByDate(ctx context.Context, dto entities.GetAvailableEquipmentByDateDto) ([]entities.Equipment, error)
		GetById(ctx context.Context, id int) (entities.Equipment, error)
		GetByEventId(ctx context.Context, id int) ([]entities.Equipment, error)
		GetUsageHistoryById(ctx context.Context, id int) ([]entities.EquipmentUsageHistory, error)

		Create(ctx context.Context, dto entities.EquipmentCreateDto) (int, error)
		Delete(ctx context.Context, id int) (int, error)
		Update(ctx context.Context, id int, dto entities.EquipmentUpdateDto) error
	}

	Event interface {
		GetAll(ctx context.Context) ([]entities.Event, error)
		GetById(ctx context.Context, id int) (entities.Event, error)
		GetByUserId(ctx context.Context, id int) ([]entities.Event, error)

		Create(ctx context.Context, dto entities.EventCreateDto) (int, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context)
	}

	Report interface {
		GetAll(ctx context.Context) ([]entities.Report, error)
		GetById(ctx context.Context, id int) (entities.Report, error)
		GetByUserId(ctx context.Context, id int) ([]entities.Report, error)

		Create(ctx context.Context, dto entities.ReportCreateDto) (int, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context)
	}

	Maintenance interface {
		GetAll(ctx context.Context) ([]entities.Maintenance, error)
		GetById(ctx context.Context, id int) (entities.Maintenance, error)

		Create(ctx context.Context, dto entities.MaintenanceCreateDto) (int, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context)
	}

	Token interface {
		GetByUserId(ctx context.Context, userId int) (entities.Token, error)
		Find(ctx context.Context, refreshToken string) (entities.Token, error)
		Save(ctx context.Context, userId int, refreshToken string) (int, error)
		Update(ctx context.Context, userId int, refreshToken string) error
		Delete(ctx context.Context, refreshToken string) error
	}

	Repository struct {
		User
		Equipment
		Event
		Report
		Maintenance
		Token
	}
)

func New(db *sqlx.DB) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		Equipment:   NewEquipmentRepository(db),
		Event:       NewEventRepository(db),
		Report:      NewReportRepository(db),
		Maintenance: NewMaintenanceRepository(db),
		Token:       NewTokenRepository(db),
	}
}
