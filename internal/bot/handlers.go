package bot

import (
	"log"
	"quiz-bot/internal/domain/models"
	"quiz-bot/internal/domain/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotState int

const (
	NormalState BotState = iota
	InTestState
	AwaitingCategoryState
)

type BotHandler struct {
	bot             *tgbotapi.BotAPI
	questionService *services.QuestionService
	currentQuestion *models.Question
	score           int
	state           BotState
	category        string
}

func NewBotHandler(bot *tgbotapi.BotAPI, questionService *services.QuestionService) *BotHandler {
	return &BotHandler{
		bot:             bot,
		questionService: questionService,
		state:           NormalState,
	}
}

func (h *BotHandler) HandleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	switch h.state {
	case NormalState:
		h.handleNormalState(message)
	case InTestState:
		h.handleInTestState(message)
	case AwaitingCategoryState:
		h.handleAwaitingCategoryState(message)
	}
}
