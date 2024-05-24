package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type (
	Auth interface {
		SignIn(user domain.UserSignInDto) (domain.User, error)
		SignOut()
	}

	User interface {
		CreateUser(dto domain.UserCreateDto) (int, error)
		GetUserById(userId int) (domain.User, error)
		GetAllUsers() ([]domain.User, error)
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
