package bot

import (
	"fmt"
	"log"
	"quiz-bot/internal/domain/models"
	"quiz-bot/internal/domain/services"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	bot             *tgbotapi.BotAPI
	questionService *services.QuestionService
	currentQuestion *models.Question
	score           uint
}

func NewBotHandler(bot *tgbotapi.BotAPI, questionService *services.QuestionService) *BotHandler {
	return &BotHandler{bot: bot, questionService: questionService}
}

func (h *BotHandler) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			h.handleMessage(update.Message)
		}
	}
}

func (h *BotHandler) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	switch message.Command() {
	case "start":
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Привет! Я бот для подготовки к собеседованиям по Go. Используйте команду /question для получения вопроса."))
	case "question":
		question := h.questionService.GetRandomQuestion()
		h.currentQuestion = &question
		questionText := question.Question
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, questionText))
	default:
		if h.currentQuestion != nil {
			userAnswer := strings.TrimSpace(message.Text)
			var responseMsg tgbotapi.MessageConfig
			if h.questionService.CheckAnswer(*h.currentQuestion, userAnswer) {
				h.score += h.currentQuestion.Points
				responseMsg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Правильно! Ваши очки: %d", h.score))

			} else {
				responseMsg = tgbotapi.NewMessage(message.Chat.ID, "Неправильно. Попробуйте ещё раз.")
			}
			h.bot.Send(responseMsg)
		}
	}
}
