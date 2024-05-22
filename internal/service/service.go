package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
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

	Service struct {
		Auth
		User
		Equipment
		Event
		Report
		Maintenance
	}
)

func New(repos *repository.Repository) *Service {
	return &Service{
		Auth:        NewAuthService(repos.Auth),
		User:        NewUserService(repos.User),
		Equipment:   NewEquipmentService(repos.Equipment),
		Event:       NewEventService(repos.Event),
		Report:      NewReportService(repos.Report),
		Maintenance: NewMaintenanceService(repos.Maintenance),
	}
}
