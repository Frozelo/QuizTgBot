package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Sender interface {
	SendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64)
	SendQuestionMessage(bot *tgbotapi.BotAPI, chatID int64, question string)
	ShowStartButtons(bot *tgbotapi.BotAPI, chatID int64)
}
