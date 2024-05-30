package service

import (
	"server-techno-flow/internal/domain"
	"server-techno-flow/internal/repository"
)

type (
	Auth interface {
		SignIn(user domain.UserSignInDto) (domain.User, string, string, error)
		SignOut(refreshToken string) error
		Refresh(refreshToken string) (domain.User, string, string, error)
	}

	User interface {
		CreateUser(dto domain.UserCreateDto) (int, error)
		GetUserById(userId int) (domain.User, error)
		GetUserByCredentials(dto domain.UserSignInDto) (domain.User, error)
		GetAllUsers() ([]domain.User, error)
		DeleteUser(id int) (int, error)
		UpdateUser(id int, dto domain.UserUpdateDto) error
	}

	Equipment interface {
		GetAllEquipment() ([]domain.Equipment, error)
		GetAvailableEquipment() ([]domain.Equipment, error)
		GetEquipmentById(id int) (domain.Equipment, error)
		GetEquipmentUsageHistoryById(id int) ([]domain.EquipmentUsageHistory, error)

		CreateEquipment(dto domain.EquipmentCreateDto) (int, error)
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
		FindRefreshToken(refreshToken string) (domain.Token, error)
		SaveRefreshToken(userId int, refreshToken string) (int, error)
		UpdateRefreshToken(userId int, refreshToken string) error
		DeleteRefreshToken(refreshToken string) error
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
	userService := NewUserService(repos.User, tokenService)
	equipmentService := NewEquipmentService(repos.Equipment)

	return &Service{
		Auth:        NewAuthService(repos.Auth, tokenService, userService),
		User:        userService,
		Equipment:   equipmentService,
		Event:       NewEventService(repos.Event, equipmentService),
		Report:      NewReportService(repos.Report),
		Maintenance: NewMaintenanceService(repos.Maintenance),
		Token:       tokenService,
	}
}
