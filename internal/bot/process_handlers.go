package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *BotHandler) startTest(chatID int64) {
	h.state = InTestState
	h.askRandomQuestion(chatID)
}

func (h *BotHandler) sendWelcomeMessage(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Привет! Я бот для подготовки к собеседованиям по Go. Используйте команду /question для получения вопроса. Используйте команду /category для выбора категории вопросов.")
	h.bot.Send(msg)
	h.showStartButtons(chatID)
}

func (h *BotHandler) promptForCategory(chatID int64) {
	h.bot.Send(tgbotapi.NewMessage(chatID, "Пожалуйста, укажите категорию вопросов."))
	h.showCategoryButtons(chatID)
	h.state = AwaitingCategoryState
}

func (h *BotHandler) askRandomQuestion(chatID int64) {
	question := h.questionService.GetRandom()
	h.currentQuestion = &question
	h.sendQuestionMessage(chatID, question.Question)
}

func (h *BotHandler) askRandomQuestionByCategory(chatID int64) {
	question, err := h.questionService.GetRandomByCategory(h.category)
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(chatID, "В данной категории нет вопросов."))
	} else {
		h.currentQuestion = &question
		h.sendQuestionMessage(chatID, question.Question)
	}
}

func (h *BotHandler) sendQuestionMessage(chatID int64, question string) {
	msg := tgbotapi.NewMessage(chatID, question)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/exit"),
		),
	)
	h.bot.Send(msg)
}

func (h *BotHandler) handleAnswer(message *tgbotapi.Message) {
	if h.currentQuestion != nil {
		userAnswer := strings.TrimSpace(message.Text)
		var responseMsg tgbotapi.MessageConfig
		if h.questionService.CheckAnswer(*h.currentQuestion, userAnswer) {
			h.score += int(h.currentQuestion.Points)
			responseMsg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Правильно! Ваши очки: %d", h.score))
		} else {
			responseMsg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Неправильно!\nПравильный ответ: %s",
				h.currentQuestion.Answer))
		}
		h.bot.Send(responseMsg)

		if h.category == "" {
			h.askRandomQuestion(message.Chat.ID)
		} else {
			h.askRandomQuestionByCategory(message.Chat.ID)
		}
	}
}

func (h *BotHandler) endTest(chatID int64) {
	finalScoreMessage := fmt.Sprintf("Тест завершен. Ваши итоговые очки: %d", h.score)
	h.bot.Send(tgbotapi.NewMessage(chatID, finalScoreMessage))
	h.reset()
}

func (h *BotHandler) reset() {
	h.currentQuestion = nil
	h.score = 0
	h.category = ""
	h.state = NormalState
}
