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
