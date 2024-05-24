package repository

import (
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/domain"
)

type (
	Auth interface {
		SignIn(dto domain.UserSignInDto) (domain.User, error)
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
		CreateEvent()
		GetEventById()
		GetAllEvents()
		DeleteEvent()
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

	Repository struct {
		User
		Auth
		Equipment
		Event
		Report
		Maintenance
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
	}
}
