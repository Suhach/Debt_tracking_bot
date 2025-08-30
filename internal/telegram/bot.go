package telegram

import (
	"os"

	"github.com/Suhach/Debt_tracking_bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		logger.Log.Error("Failed to init bot", zap.Error(err))
		return nil
	}

	logger.Log.Info("Bot initialized successfully", zap.String("bot_username", bot.Self.UserName))

	bot.Debug = true
	return bot
}
