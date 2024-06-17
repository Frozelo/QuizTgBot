package bot

import (
	"fmt"
	"log"
	"quiz-bot/internal/domain/models"
	"quiz-bot/internal/domain/services"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotState int

const (
	NormalState BotState = iota
	AwaitingCategoryState
)

type BotHandler struct {
	bot             *tgbotapi.BotAPI
	questionService *services.QuestionService
	currentQuestion *models.Question
	score           uint
	state           int
}

func NewBotHandler(bot *tgbotapi.BotAPI, questionService *services.QuestionService) *BotHandler {
	return &BotHandler{bot: bot, questionService: questionService, state: int(NormalState)}
}

func (h *BotHandler) handleMessage(message *tgbotapi.Message) {

	switch h.state {
	case int(NormalState):
		h.handleNormalState(message)
	case int(AwaitingCategoryState):
		h.handleAwaitingCategoryState(message)
	}
}

func (h *BotHandler) handleNormalState(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Привет! Я бот для подготовки к собеседованиям по Go. Используйте команду /question для получения вопроса. Используйте команду /category для выбора категории вопросов."))
	case "question":
		question := h.questionService.GetRandom()
		h.currentQuestion = &question
		questionText := question.Question
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, questionText))
	case "category":
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, укажите категорию вопросов."))
		h.state = int(AwaitingCategoryState)
	default:
		h.handleAnswer(message)
	}
}

func (h *BotHandler) handleAwaitingCategoryState(message *tgbotapi.Message) {
	category := strings.TrimSpace(message.Text)
	question, err := h.questionService.GetRandomByCategory(category)
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Категория вопросов не найдена."))
		h.state = int(NormalState)
		return
	}

	h.currentQuestion = &question
	questionText := question.Question
	h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, questionText))
	h.state = int(NormalState)
}

func (h *BotHandler) handleAnswer(message *tgbotapi.Message) {
	if h.currentQuestion != nil {
		log.Printf("Total current question is %v", h.currentQuestion)
		userAnswer := strings.TrimSpace(message.Text)
		var responseMsg tgbotapi.MessageConfig
		if h.questionService.CheckAnswer(*h.currentQuestion, userAnswer) {
			h.score += h.currentQuestion.Points
			responseMsg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Правильно! Ваши очки: %d", h.score))
		} else {
			responseMsg = tgbotapi.NewMessage(message.Chat.ID, "Неправильно. Попробуйте ещё раз.")
		}
		h.bot.Send(responseMsg)

		question := h.questionService.GetRandom()
		h.currentQuestion = &question
		questionText := question.Question
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, questionText))
	}
}
