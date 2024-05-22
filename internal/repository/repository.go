package repository

import (
	"github.com/jmoiron/sqlx"
	"server-techno-flow/internal/domain"
)

type (
	Auth interface {
		SignUp()
		SignIn()
		SignOut()
	}

	User interface {
		CreateUser(user domain.User) (int, error)
		GetUserByUsername(username string) (domain.User, error)
		GetAllUsers() ([]domain.User, error)
		DeleteUser(id int) (int, error)
		UpdateUser(id int, userDto domain.UserUpdateDto) error
	}

	Equipment interface {
		CreateEquipment()
		GetEquipment()
		GetAllEquipment()
		DeleteEquipment()
		UpdateEquipment()
	}

	Event interface {
		CreateEvent()
		GetEvent()
		GetAllEvents()
		DeleteEvent()
		UpdateEvent()
	}

	Report interface {
		CreateReport()
		GetReport()
		GetAllReports()
		DeleteReport()
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
