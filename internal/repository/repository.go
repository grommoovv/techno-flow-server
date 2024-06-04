package repository

import (
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/entities"
)

type (
	Auth interface {
		SignIn()
		SignOut()
	}

	User interface {
		GetAll() ([]entities.User, error)
		GetById(id int) (entities.User, error)
		GetByCredentials(dto entities.UserSignInDto) (entities.User, error)

		Create(dto entities.UserCreateDto) (int, error)
		Delete(id int) (int, error)
		Update(id int, dto entities.UserUpdateDto) error
	}

	Equipment interface {
		GetAll() ([]entities.Equipment, error)
		GetAvailableByDate(dto entities.GetAvailableEquipmentByDateDto) ([]entities.Equipment, error)
		GetById(id int) (entities.Equipment, error)
		GetByEventId(id int) ([]entities.Equipment, error)
		GetUsageHistoryById(id int) ([]entities.EquipmentUsageHistory, error)

		Create(dto entities.EquipmentCreateDto) (int, error)
		Delete(id int) (int, error)
		Update(id int, dto entities.EquipmentUpdateDto) error
	}

	Event interface {
		GetAll() ([]entities.Event, error)
		GetById(id int) (entities.Event, error)
		GetByUserId(id int) ([]entities.Event, error)

		Create(dto entities.EventCreateDto) (int, error)
		Delete(id int) error
		Update()
	}

	Report interface {
		GetAllReports() ([]entities.Report, error)
		GetReportById(id int) (entities.Report, error)
		CreateReport(dto entities.ReportCreateDto) (int, error)
		DeleteReport(id int) error
		UpdateReport()
	}

	Maintenance interface {
		GetAll() ([]entities.Maintenance, error)
		GetById(id int) (entities.Maintenance, error)

		Create(dto entities.MaintenanceCreateDto) (int, error)
		Delete(id int) error
		Update()
	}

	Token interface {
		GetTokenByUserId(userId int) (entities.Token, error)
		FindRefreshToken(refreshToken string) (entities.Token, error)
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
