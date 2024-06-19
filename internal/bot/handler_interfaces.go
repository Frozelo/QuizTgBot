package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandlerInterface interface {
	HandleMessage(message *tgbotapi.Message)
	HandleCallback(callback *tgbotapi.CallbackQuery)
}
