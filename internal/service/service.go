package service

import (
	"math/rand"
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type (
	Auth interface {
		SignIn(user entities.UserSignInDto) (entities.User, string, string, error)
		SignOut(refreshToken string) error
		Refresh(refreshToken string) (entities.User, string, string, error)
	}

	User interface {
		CreateUser(dto entities.UserCreateDto) (int, error)
		GetUserById(userId int) (entities.User, error)
		GetUserByCredentials(dto entities.UserSignInDto) (entities.User, error)
		GetAllUsers() ([]entities.User, error)
		DeleteUser(id int) (int, error)
		UpdateUser(id int, dto entities.UserUpdateDto) error
	}

	Equipment interface {
		GetAllEquipment() ([]entities.Equipment, error)
		GetAvailableEquipmentByDate(dto entities.GetAvailableEquipmentByDateDto) ([]entities.Equipment, error)
		GetEquipmentById(id int) (entities.Equipment, error)
		GetEquipmentByEventId(id int) ([]entities.Equipment, error)
		GetEquipmentUsageHistoryById(id int) ([]entities.EquipmentUsageHistory, error)

		CreateEquipment(dto entities.EquipmentCreateDto) (int, error)
		DeleteEquipment(id int) (int, error)
		UpdateEquipment(id int, dto entities.EquipmentUpdateDto) error
	}

	Event interface {
		CreateEvent(dto entities.EventCreateDto) (int, error)
		GetAllEvents() ([]entities.Event, error)
		GetEventById(id int) (entities.Event, error)
		GetEventsByUserId(id int) ([]entities.Event, error)
		DeleteEvent(id int) error
		UpdateEvent()
	}

	Report interface {
		GetAllReports() ([]entities.Report, error)
		GetReportById(id int) (entities.Report, error)
		GetReportsByUserId(id int) ([]entities.Report, error)

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
		NewRefreshToken(userId int, username string) (string, error)
		NewAccessToken(userId int, username string) (string, error)
		ParseRefreshToken(accessToken string) (int, error)
		ParseAccessToken(accessToken string) (int, error)
		GetTokenByUserId(userId int) (entities.Token, error)
		FindRefreshToken(refreshToken string) (entities.Token, error)
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

func New(repos *repository.Repository, r *rand.Rand) *Service {
	tokenService := NewTokenService(repos.Token)
	userService := NewUserService(repos.User, tokenService)
	equipmentService := NewEquipmentService(repos.Equipment)
	maintenanceService := NewMaintenanceService(repos.Maintenance, r)

	return &Service{
		Auth:        NewAuthService(tokenService, userService),
		User:        userService,
		Equipment:   equipmentService,
		Event:       NewEventService(repos.Event, equipmentService),
		Report:      NewReportService(repos.Report, equipmentService, maintenanceService),
		Maintenance: maintenanceService,
		Token:       tokenService,
	}
}
