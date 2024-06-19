package bot

import (
	"quiz-bot/internal/domain/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sender interface {
	SendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64)
	SendQuestionMessage(bot *tgbotapi.BotAPI, chatID int64, question models.Question)
	ShowStartButtons(bot *tgbotapi.BotAPI, chatID int64)
}
