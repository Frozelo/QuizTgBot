package bot

import (
	"fmt"
	"quiz-bot/internal/domain/models"
	"quiz-bot/internal/domain/services"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StateContext struct {
	currentQuestion *models.Question
	score           int
	questionService *services.QuestionService
	messageSender   Sender
	stateHandler    Stater
}

func (ctx *StateContext) Reset() {
	ctx.currentQuestion = nil
	ctx.score = 0
}

type NormalStateHanlder struct {
	ctx *StateContext
}

func NewNormalStateHandler(ctx *StateContext) *NormalStateHanlder {
	return &NormalStateHanlder{ctx: ctx}
}

func (h *NormalStateHanlder) Handle(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		h.ctx.messageSender.SendWelcomeMessage(bot, message.Chat.ID)

	case "question", "test":
		h.startTest(bot, message.Chat.ID)

	// case "category":
	// 	h.askForCategoryCallback(bot, message.Chat.ID)

	case "exit":
		h.endTest(bot, message.Chat.ID)

	default:
		h.ctx.messageSender.SendErrorMessage(bot, message.Chat.ID)
	}
}

func (h *NormalStateHanlder) HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	h.ctx.stateHandler = NewTestStateHandler(h.ctx)
	h.ctx.stateHandler.HandleCallback(bot, callback)
}

func (h *NormalStateHanlder) startTest(bot *tgbotapi.BotAPI, chatID int64) {
	h.ctx.stateHandler = NewTestStateHandler(h.ctx)
	h.ctx.stateHandler.Handle(bot, &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: chatID}, Text: "/question"})
}

func (h *NormalStateHanlder) endTest(bot *tgbotapi.BotAPI, chatID int64) {
	finalScoreMessage := fmt.Sprintf("Тест завершен. Ваши итоговые очки: %d", h.ctx.score)
	bot.Send(tgbotapi.NewMessage(chatID, finalScoreMessage))
	h.ctx.Reset()
	h.ctx.stateHandler = NewNormalStateHandler(h.ctx)
}

type TestStateHandler struct {
	ctx *StateContext
}

func NewTestStateHandler(ctx *StateContext) *TestStateHandler {
	return &TestStateHandler{ctx: ctx}
}

func (h *TestStateHandler) Handle(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message.Command() == "exit" {
		h.endTest(bot, message.Chat.ID)
		return
	}
	h.askRandomQuestion(bot, message.Chat.ID)
}

func (h *TestStateHandler) HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	userAnswer := strings.TrimSpace(callback.Data)
	responseMsg := h.generateResponseMessage(userAnswer, callback.Message.Chat.ID, callback.Message.MessageID)
	bot.Send(responseMsg)
	h.askRandomQuestion(bot, callback.Message.Chat.ID)
}

func (h *TestStateHandler) generateResponseMessage(userAnswer string, chatID int64, messageID int) tgbotapi.EditMessageTextConfig {
	var responseMsg tgbotapi.EditMessageTextConfig
	correct, err := h.ctx.questionService.CheckUserAnswer(h.ctx.currentQuestion, userAnswer)
	if err != nil {
		responseMsg = tgbotapi.NewEditMessageText(chatID, messageID, fmt.Sprintf("Произошла ошибка при проверке ответа: %v", err))
	} else {
		if correct {
			responseMsg = tgbotapi.NewEditMessageText(chatID, messageID, fmt.Sprintf("Правильно! Ваши очки: %d", h.handleCorrectAnswer()))
		} else {
			answer := h.ctx.questionService.GetRightAnswerText(h.ctx.currentQuestion)
			responseMsg = tgbotapi.NewEditMessageText(chatID, messageID, fmt.Sprintf("Неправильно!\nПравильный ответ: %s", answer))
		}
	}
	return responseMsg
}

func (h *TestStateHandler) handleCorrectAnswer() int {
	h.ctx.score += int(h.ctx.currentQuestion.Points)
	return h.ctx.score
}

func (h *TestStateHandler) askRandomQuestion(bot *tgbotapi.BotAPI, chatID int64) {
	question := h.ctx.questionService.GetRandom()
	h.ctx.currentQuestion = &question
	h.ctx.messageSender.SendQuestionMessage(bot, chatID, question)
}

func (h *TestStateHandler) endTest(bot *tgbotapi.BotAPI, chatID int64) {
	finalScoreMessage := fmt.Sprintf("Тест завершен. Ваши итоговые очки: %d", h.ctx.score)
	bot.Send(tgbotapi.NewMessage(chatID, finalScoreMessage))
	h.ctx.Reset()
	h.ctx.stateHandler = NewNormalStateHandler(h.ctx)
}
