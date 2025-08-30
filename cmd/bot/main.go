package main

import (
	"context"
	"log"

	"github.com/Suhach/Debt_tracking_bot/internal/adapters"
	"github.com/Suhach/Debt_tracking_bot/internal/database/postgres"
	"github.com/Suhach/Debt_tracking_bot/internal/repository"
	"github.com/Suhach/Debt_tracking_bot/internal/service"
	"github.com/Suhach/Debt_tracking_bot/internal/telegram"
	"github.com/Suhach/Debt_tracking_bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Error("Application panicked", zap.Any("error", err))
		}
	}()

	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("Error loading .env file")
	}

	LoggerError := logger.Init()
	if LoggerError != nil {
		log.Fatal("Failed to initialize logger")
	}
	defer logger.Log.Sync()
	logger.Log.Info("Logger initialized successfully")

	postgres.InitPGX()

	bot := telegram.Init()

	repoDebt := repository.NewDebtRepo(postgres.Pool)
	repoUser := repository.NewUser(postgres.Pool)

	serviceDebt := service.NewDebtService(repoDebt)
	serviceUser := service.NewUserService(repoUser)

	handlerDebt := adapters.NewDebtAdapter(serviceDebt)
	handlerUser := adapters.NewUser(serviceUser)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for update := range updates {
		if update.Message != nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "/newdebt":
			handlerDebt.CreateDebt(ctx, update.Message)
		case "/newuser":
			handlerUser.NewUser()
		default:
			msg.Text = "I don't know that command"
		}
		if _, err := bot.Send(msg); err != nil {
			logger.Log.Fatal("Failed to send message", zap.Error(err))
		}
	}
}
