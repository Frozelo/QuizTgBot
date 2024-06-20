package bot

import (
	"quiz-bot/internal/domain/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	bot             *tgbotapi.BotAPI
	questionService *services.QuestionService
	stateHandler    Stater
	messageSender   Sender
}

func NewBotHandler(bot *tgbotapi.BotAPI, questionService *services.QuestionService,
	stateHandler Stater, messageSender Sender) *BotHandler {
	return &BotHandler{
		bot:             bot,
		questionService: questionService,
		stateHandler:    stateHandler,
		messageSender:   messageSender,
	}
}

func (h *BotHandler) HandleMessage(message *tgbotapi.Message) {
	h.stateHandler.HandleState(h.bot, message)
}

func (h *BotHandler) HandleCallback(callback *tgbotapi.CallbackQuery) {
	h.stateHandler.HandleCallback(h.bot, callback)
}
