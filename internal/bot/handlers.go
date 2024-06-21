package bot

import (
	"quiz-bot/internal/domain/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	botAPI       *tgbotapi.BotAPI
	stateHandler Stater
	stateContext *StateContext
}

func NewBot(botAPI *tgbotapi.BotAPI, questionService *services.QuestionService, messageSender Sender) *BotHandler {
	stateContext := &StateContext{
		currentQuestion: nil,
		score:           0,
		questionService: questionService,
		messageSender:   messageSender,
	}
	botHandler := &BotHandler{
		stateContext: stateContext,
		botAPI:       botAPI,
	}
	botHandler.stateHandler = NewNormalStateHandler(stateContext)
	return botHandler
}

func (h *BotHandler) HandleMessage(message *tgbotapi.Message) {
	h.stateHandler.Handle(h.botAPI, message)
}

func (h *BotHandler) HandleCallback(callback *tgbotapi.CallbackQuery) {
	h.stateHandler.HandleCallback(h.botAPI, callback)
}
