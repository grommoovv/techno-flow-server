package service

import (
	"context"
	"math/rand"
	"server-techno-flow/internal/entities"
	"server-techno-flow/internal/repository"
)

type (
	Auth interface {
		SignIn(ctx context.Context, user entities.UserSignInDto) (entities.User, string, string, error)
		SignOut(ctx context.Context, refreshToken string) error
		Refresh(ctx context.Context, refreshToken string) (entities.User, string, string, error)
	}

	User interface {
		CreateUser(ctx context.Context, dto entities.UserCreateDto) (int, error)
		GetUserById(ctx context.Context, userId int) (entities.User, error)
		GetUserByCredentials(ctx context.Context, dto entities.UserSignInDto) (entities.User, error)
		GetAllUsers(ctx context.Context) ([]entities.User, error)
		DeleteUser(ctx context.Context, id int) (int, error)
		UpdateUser(ctx context.Context, id int, dto entities.UserUpdateDto) error
	}

	Equipment interface {
		GetAllEquipment(ctx context.Context) ([]entities.Equipment, error)
		GetAvailableEquipmentByDate(ctx context.Context, dto entities.GetAvailableEquipmentByDateDto) ([]entities.Equipment, error)
		GetEquipmentById(ctx context.Context, id int) (entities.Equipment, error)
		GetEquipmentByEventId(ctx context.Context, id int) ([]entities.Equipment, error)
		GetEquipmentUsageHistoryById(ctx context.Context, id int) ([]entities.EquipmentUsageHistory, error)

		CreateEquipment(ctx context.Context, dto entities.EquipmentCreateDto) (int, error)
		DeleteEquipment(ctx context.Context, id int) (int, error)
		UpdateEquipment(ctx context.Context, id int, dto entities.EquipmentUpdateDto) error
	}

	Event interface {
		CreateEvent(ctx context.Context, dto entities.EventCreateDto) (int, error)
		GetAllEvents(ctx context.Context) ([]entities.Event, error)
		GetEventById(ctx context.Context, id int) (entities.Event, error)
		GetEventsByUserId(ctx context.Context, id int) ([]entities.Event, error)
		DeleteEvent(ctx context.Context, id int) error
		UpdateEvent(ctx context.Context)
	}

	Report interface {
		GetAllReports(ctx context.Context) ([]entities.Report, error)
		GetReportById(ctx context.Context, id int) (entities.Report, error)
		GetReportsByUserId(ctx context.Context, id int) ([]entities.Report, error)

		CreateReport(ctx context.Context, dto entities.ReportCreateDto) (int, error)
		DeleteReport(ctx context.Context, id int) error
		UpdateReport(ctx context.Context)
	}

	Maintenance interface {
		GetAll(ctx context.Context) ([]entities.Maintenance, error)
		GetById(ctx context.Context, id int) (entities.Maintenance, error)

		Create(ctx context.Context, dto entities.MaintenanceCreateDto) (int, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context)
	}

	Token interface {
		NewRefreshToken(ctx context.Context, userId int, username string) (string, error)
		NewAccessToken(ctx context.Context, userId int, username string) (string, error)
		ParseRefreshToken(ctx context.Context, accessToken string) (int, error)
		ParseAccessToken(ctx context.Context, accessToken string) (int, error)
		GetTokenByUserId(ctx context.Context, userId int) (entities.Token, error)
		FindRefreshToken(ctx context.Context, refreshToken string) (entities.Token, error)
		SaveRefreshToken(ctx context.Context, userId int, refreshToken string) (int, error)
		UpdateRefreshToken(ctx context.Context, userId int, refreshToken string) error
		DeleteRefreshToken(ctx context.Context, refreshToken string) error
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
