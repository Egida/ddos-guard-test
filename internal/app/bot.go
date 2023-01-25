package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Shteyd/ddos-guard-test/config"
	"github.com/Shteyd/ddos-guard-test/internal/controller/tgbot"
	"github.com/Shteyd/ddos-guard-test/internal/infrastructure/mathservice"
	"github.com/Shteyd/ddos-guard-test/internal/infrastructure/repository"
	"github.com/Shteyd/ddos-guard-test/internal/usecase"
	"github.com/Shteyd/ddos-guard-test/pkg/logger"
	"github.com/Shteyd/ddos-guard-test/pkg/postgres"
)

func RunBot(cfg *config.Config) {
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
	mathUseCase := usecase.NewMathUC(mathservice.New())
	// init tg bot
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := tgbot.InitBot(cfg, logger, userUseCase, mathUseCase)
	if err != nil {
		logger.Fatal(fmt.Errorf("bot - Run - tgbot.InitBot: %w", err))
	}

	b.Start(ctx)
}
