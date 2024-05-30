package repository

import (
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/domain"
)

type (
	Auth interface {
		GetUserByCredentials(dto domain.UserSignInDto) (domain.User, error)
		SignOut()
	}

	User interface {
		GetAll() ([]domain.User, error)
		GetById(id int) (domain.User, error)
		Create(dto domain.UserCreateDto) (int, error)
		Delete(id int) (int, error)
		Update(id int, dto domain.UserUpdateDto) error
	}

	Equipment interface {
		GetAll() ([]domain.Equipment, error)
		GetAvailable() ([]domain.Equipment, error)
		GetById(id int) (domain.Equipment, error)
		GetUsageHistoryById(id int) ([]domain.EquipmentUsageHistory, error)

		Create(dto domain.EquipmentCreateDto) (int, error)
		Delete(id int) (int, error)
		Update(id int, dto domain.EquipmentUpdateDto) error
	}

	Event interface {
		GetAll() ([]domain.Event, error)
		GetById(id int) (domain.Event, error)
		Create(dto domain.EventCreateDto) (int, error)
		Delete(id int) (int, error)
		Update()
	}

	Report interface {
		CreateReport(dto domain.ReportCreateDto) (int, error)
		GetAllReports() ([]domain.Report, error)
		GetReportById(id int) (domain.Report, error)
		DeleteReport(id int) (int, error)
		UpdateReport()
	}

	Maintenance interface {
		GetAll()
		GetById()

		Create()
		Delete()
		Update()
	}

	Token interface {
		GetTokenByUserId(userId int) (domain.Token, error)
		FindRefreshToken(refreshToken string) (domain.Token, error)
		SaveRefreshToken(userId int, refreshToken string) (int, error)
		UpdateToken(userId int, refreshToken string) error
		DeleteToken(refreshToken string) error
	}

	Repository struct {
		User
		Auth
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
		Auth:        NewAuthRepository(db),
		Equipment:   NewEquipmentRepository(db),
		Event:       NewEventRepository(db),
		Report:      NewReportRepository(db),
		Maintenance: NewMaintenanceRepository(db),
		Token:       NewTokenRepository(db),
	}
}
