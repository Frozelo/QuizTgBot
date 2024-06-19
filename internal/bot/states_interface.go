package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Stater interface {
	HandleState(bot *tgbotapi.BotAPI, message *tgbotapi.Message)
	HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery)
}
