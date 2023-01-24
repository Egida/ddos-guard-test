package tgbot

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shteyd/ddos-guard-test/internal/usecase"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update, uc usecase.Math) {
	if update.Message == nil {
		return
	}

	text := update.Message.Text
	if strings.Contains(text, "/start") {
		sendStart(ctx, b, update)
	} else if strings.Contains(text, "/calculate") {
		sendResult(ctx, b, update, uc)
	}
}

func sendStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет! Я бот калькулятор. Если хочешь получить ответ на свой пример, то напиши /calculate и формулу (пример: /calculate 2 + 2)",
	})
}

func sendResult(ctx context.Context, b *bot.Bot, update *models.Update, m usecase.Math) {
	text := validateCalculate(update.Message.Text)
	number, err := m.Calculate(text)
	if err != nil {
		sendErrorMessage(ctx, b, update)
		return
	}
	result := validateMathResult(number)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Ответ: %s", result),
	})
}
