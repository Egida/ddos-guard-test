package tgbot

import (
	"fmt"
	"context"

	"github.com/Shteyd/ddos-guard-test/config"
	"github.com/Shteyd/ddos-guard-test/internal/usecase"
	"github.com/Shteyd/ddos-guard-test/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func InitBot(cfg *config.Config, logger *logger.Logger, userUC usecase.User, mathUC usecase.Math) (*bot.Bot, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(func(ctx context.Context, bot *bot.Bot, update *models.Update) {
			handler(ctx, bot, update, mathUC)
		}),
		bot.WithMiddlewares(func(next bot.HandlerFunc) bot.HandlerFunc {
			return func(ctx context.Context, b *bot.Bot, update *models.Update) {
				if update.Message != nil {
					username := update.Message.From.Username
					id, err := userUC.GetUserID(username)
					if err != nil {
						if err := userUC.Store(username); err != nil {
							logger.Info(fmt.Errorf("bot - middleware - userUC.Store: %w", err))
							sendErrorMessage(ctx, b, update)
							return
						}
					}

					if id == 0 {
						if err := userUC.Store(username); err != nil {
							logger.Info(fmt.Errorf("bot - middleware - userUC.Store: %w", err))
							sendErrorMessage(ctx, b, update)
							return
						}
					}
				}
				next(ctx, b, update)
			}
		}),
	}

	return bot.New(cfg.Bot.Token, opts...)
}
