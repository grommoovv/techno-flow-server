package app

import (
	"context"
	"os"
	"os/signal"
	"server-techno-flow/internal/config"
	"server-techno-flow/internal/database/postgres"
	"server-techno-flow/internal/handler"
	"server-techno-flow/internal/repository"
	"server-techno-flow/internal/server"
	"server-techno-flow/internal/service"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Run() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	conf, err := config.Init()

	if err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
		return
	}

	ps, err := postgres.NewPostgres(postgres.Postgres{
		Host:     conf.Postgres.Host,
		Port:     conf.Postgres.Port,
		Username: conf.Postgres.Username,
		Password: conf.Postgres.Password,
		DBName:   conf.Postgres.Dbname,
		SSLMode:  conf.Postgres.SslMode,
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
		return
	}

	repos := repository.New(ps)
	services := service.New(repos)
	handlers := handler.New(services)

	srv := server.New(conf, handlers.Init())

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Infof("server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("error occured while stopping http server: %s", err.Error())
	}

	if err := ps.Close(); err != nil {
		logrus.Fatalf("error occured while closing postgres: %s", err.Error())
	}
}
