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
	case AwaitingCategoryState:
		h.handleAwaitingCategoryState(message)
	}
}

func (h *BotHandler) handleNormalState(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Привет! Я бот для подготовки к собеседованиям по Go. Используйте команду /question для получения вопроса. Используйте команду /category для выбора категории вопросов."))
	case "question":
		h.askRandomQuestion(message.Chat.ID)
	case "category":
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, укажите категорию вопросов."))
		h.showCategoryButtons(message.Chat.ID)
	case "exit":
		h.endTest(message.Chat.ID)
	default:
		h.handleAnswer(message)
	}
}

func (h *BotHandler) handleAwaitingCategoryState(message *tgbotapi.Message) {
	h.category = strings.TrimSpace(message.Text)
	h.state = NormalState
	h.askRandomQuestionByCategory(message.Chat.ID)
}

func (h *BotHandler) handleAnswer(message *tgbotapi.Message) {
	if h.currentQuestion != nil {
		userAnswer := strings.TrimSpace(message.Text)
		var responseMsg tgbotapi.MessageConfig
		if h.questionService.CheckAnswer(*h.currentQuestion, userAnswer) {
			h.score += int(h.currentQuestion.Points)
			responseMsg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Правильно! Ваши очки: %d", h.score))
		} else {
			responseMsg = tgbotapi.NewMessage(message.Chat.ID, "Неправильно. Попробуйте ещё раз.")
		}
		h.bot.Send(responseMsg)

		if h.category == "" {
			h.askRandomQuestion(message.Chat.ID)
		} else {
			h.askRandomQuestionByCategory(message.Chat.ID)
		}
	}
}

func (h *BotHandler) askRandomQuestion(chatID int64) {
	question := h.questionService.GetRandom()
	h.currentQuestion = &question
	h.bot.Send(tgbotapi.NewMessage(chatID, question.Question))
}

func (h *BotHandler) askRandomQuestionByCategory(chatID int64) {
	question, err := h.questionService.GetRandomByCategory(h.category)
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(chatID, "В данной категории нет вопросов."))
	} else {
		h.currentQuestion = &question
		h.bot.Send(tgbotapi.NewMessage(chatID, question.Question))
	}
}

func (h *BotHandler) showCategoryButtons(chatID int64) {
	categories, err := h.questionService.GetCategories()
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(chatID, "В данной категории нет вопросов."))
		return
	}
	var buttons []tgbotapi.KeyboardButton
	for _, category := range categories {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(category))

	}

	replyMarkup := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(buttons...))
	msg := tgbotapi.NewMessage(chatID, "Выберите категорию:")
	msg.ReplyMarkup = replyMarkup

	h.bot.Send(msg)
	h.state = AwaitingCategoryState
}

func (h *BotHandler) endTest(chatID int64) {
	finalScoreMessage := fmt.Sprintf("Тест завершен. Ваши итоговые очки: %d", h.score)
	h.bot.Send(tgbotapi.NewMessage(chatID, finalScoreMessage))
	h.currentQuestion = nil
	h.score = 0
	h.category = ""
	h.state = NormalState
}
