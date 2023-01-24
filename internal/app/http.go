package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shteyd/ddos-guard-test/config"
	v1 "github.com/Shteyd/ddos-guard-test/internal/controller/http/v1"
	"github.com/Shteyd/ddos-guard-test/internal/infrastructure/repository"
	"github.com/Shteyd/ddos-guard-test/internal/usecase"
	"github.com/Shteyd/ddos-guard-test/pkg/httpserver"
	"github.com/Shteyd/ddos-guard-test/pkg/logger"
	"github.com/Shteyd/ddos-guard-test/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func RunHTTP(cfg *config.Config) {
	// init logger
	logger := logger.New(cfg.Log.Level)

	// init postgres database
	db, err := postgres.NewPostgres(postgres.Config{
		Host:     cfg.PG.Host,
		Port:     cfg.PG.Port,
		Username: cfg.PG.Username,
		Password: cfg.PG.Password,
		DBName:   cfg.PG.DBName,
		SSLMode:  cfg.PG.SSLMode,
	})
	if err != nil {
		logger.Fatal(fmt.Errorf("app - Run - postgres.NewPostgres: %w", err))
	}
	defer db.Close()

	// init usecases
	userUseCase := usecase.NewUsersUC(repository.New(db))

	// init HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, userUseCase, logger)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	if err := httpServer.Shutdown(); err != nil {
		logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
