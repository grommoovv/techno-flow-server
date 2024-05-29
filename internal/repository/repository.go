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
		CreateUser(dto domain.UserCreateDto) (int, error)
		GetAllUsers() ([]domain.User, error)
		GetUserById(id int) (domain.User, error)
		DeleteUser(id int) (int, error)
		UpdateUser(id int, dto domain.UserUpdateDto) error
	}

	Equipment interface {
		CreateEquipment(dto domain.EquipmentCreateDto) (int, error)
		GetAllEquipment() ([]domain.Equipment, error)
		GetEquipmentById(id int) (domain.Equipment, error)
		DeleteEquipment(id int) (int, error)
		UpdateEquipment(id int, dto domain.EquipmentUpdateDto) error
	}

	Event interface {
		CreateEvent(dto domain.EventCreateDto) (int, error)
		GetAllEvents() ([]domain.Event, error)
		GetEventById(id int) (domain.Event, error)
		DeleteEvent(id int) (int, error)
		UpdateEvent()
	}

	Report interface {
		CreateReport(dto domain.ReportCreateDto) (int, error)
		GetAllReports() ([]domain.Report, error)
		GetReportById(id int) (domain.Report, error)
		DeleteReport(id int) (int, error)
		UpdateReport()
	}

	Maintenance interface {
		CreateMaintenance()
		GetMaintenance()
		GetAllMaintenance()
		DeleteMaintenance()
		UpdateMaintenance()
	}

	Token interface {
		GetTokenByUserId(userId int) (domain.Token, error)
		FindToken(refreshToken string) (domain.Token, error)
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
