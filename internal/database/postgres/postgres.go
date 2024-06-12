package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	UsersTable          = "users"
	EquipmentTable      = "equipment"
	EventsTable         = "events"
	EquipmentUsageTable = "equipment_usage"
	ReportsTable        = "reports"
	MaintenanceTable    = "maintenance"
	TokensTable         = "tokens"
	RolesTable          = "roles"
	UserRolesTable      = "user_roles"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func New(cfg Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	logrus.Infof("postgres connection successfully established")

	return db, nil
}
