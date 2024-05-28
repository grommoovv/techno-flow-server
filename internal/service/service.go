package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type (
	Auth interface {
		SignIn(user domain.UserSignInDto) (domain.User, string, string, error)
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

	Token interface {
		NewRefreshToken(userId int, username string) (string, error)
		NewAccessToken(userId int, username string) (string, error)
		ParseRefreshToken(accessToken string) (int, error)
		ParseAccessToken(accessToken string) (int, error)
		GetTokenByUserId(userId int) (domain.Token, error)
		SaveToken(userId int, refreshToken string) (int, error)
		UpdateToken(userId int, refreshToken string) (int, error)
		DeleteToken(refreshToken string) error
	}

	Service struct {
		Auth
		User
		Equipment
		Event
		Report
		Maintenance
		Token
	}
)

func New(repos *repository.Repository) *Service {
	tokenService := NewTokenService(repos.Token)
	return &Service{
		Auth:        NewAuthService(repos.Auth, tokenService),
		User:        NewUserService(repos.User, tokenService),
		Equipment:   NewEquipmentService(repos.Equipment),
		Event:       NewEventService(repos.Event),
		Report:      NewReportService(repos.Report),
		Maintenance: NewMaintenanceService(repos.Maintenance),
		Token:       tokenService,
	}
}
