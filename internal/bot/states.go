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
	InTestState
)

type StateHandler struct {
	state           BotState
	currentQuestion *models.Question
	score           int
	category        string
	questionService *services.QuestionService
	messageSender   Sender
}

func NewStateHandler(questionService *services.QuestionService, messageSender Sender) *StateHandler {
	return &StateHandler{
		state:           NormalState,
		questionService: questionService,
		messageSender:   messageSender,
	}
}

func (h *StateHandler) HandleState(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	switch h.state {
	case NormalState:
		h.handleDefaultState(bot, message)
	case InTestState:
		h.handleInTestState(bot, message)
	}
}

func (h *StateHandler) HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	switch h.state {
	case InTestState:
		h.handleAnswerCallback(bot, callback)
	}
}

func (h *StateHandler) handleDefaultState(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		h.messageSender.SendWelcomeMessage(bot, message.Chat.ID)
	case "question", "test":
		h.startTest(bot, message.Chat.ID)
	case "exit":
		h.endTest(bot, message.Chat.ID)
	default:
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Используйте /start, /question или /exit."))
	}
}

func (h *StateHandler) handleInTestState(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message.Command() == "exit" {
		h.endTest(bot, message.Chat.ID)
	}
}

func (h *StateHandler) startTest(bot *tgbotapi.BotAPI, chatID int64) {
	h.state = InTestState
	h.askRandomQuestion(bot, chatID)
}

func (h *StateHandler) askRandomQuestion(bot *tgbotapi.BotAPI, chatID int64) {
	question := h.questionService.GetRandom()
	h.currentQuestion = &question
	h.messageSender.SendQuestionMessage(bot, chatID, question)
}

func (h *StateHandler) handleAnswerCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	userAnswer := strings.TrimSpace(callback.Data)
	var responseMsg tgbotapi.EditMessageTextConfig
	correct, err := h.questionService.CheckAnswer(*h.currentQuestion, userAnswer)
	if err != nil {
		responseMsg = tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, fmt.Sprintf("Произошла ошибка при проверке ответа: %v", err))
	} else {
		if correct {
			h.score += int(h.currentQuestion.Points)
			responseMsg = tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, fmt.Sprintf("Правильно! Ваши очки: %d", h.score))
		} else {
			responseMsg = tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, fmt.Sprintf("Неправильно!\nПравильный ответ: %v", h.currentQuestion.RightAnswerID))
		}
	}

	bot.Send(responseMsg)

	if h.category == "" {
		h.askRandomQuestion(bot, callback.Message.Chat.ID)
	}
}

func (h *StateHandler) endTest(bot *tgbotapi.BotAPI, chatID int64) {
	finalScoreMessage := fmt.Sprintf("Тест завершен. Ваши итоговые очки: %d", h.score)
	bot.Send(tgbotapi.NewMessage(chatID, finalScoreMessage))
	h.reset()
}

func (h *StateHandler) reset() {
	h.currentQuestion = nil
	h.score = 0
	h.category = ""
	h.state = NormalState
}
