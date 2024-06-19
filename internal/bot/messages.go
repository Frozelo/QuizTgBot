package bot

import (
	"quiz-bot/internal/domain/models"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageSender struct{}

func NewMessageSender() *MessageSender {
	return &MessageSender{}
}

func (s *MessageSender) SendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Привет! Я бот для подготовки к собеседованиям по Go. Используйте команду /question для получения вопроса. Используйте команду /category для выбора категории вопросов.")
	bot.Send(msg)
	s.ShowStartButtons(bot, chatID)
}

func (s *MessageSender) SendQuestionMessage(bot *tgbotapi.BotAPI, chatID int64, question models.Question) {
	msg := tgbotapi.NewMessage(chatID, question.Question)
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, answer := range question.Answers {
		button := tgbotapi.NewInlineKeyboardButtonData(answer.Text, strconv.Itoa(answer.ID))
		row := tgbotapi.NewInlineKeyboardRow(button)
		rows = append(rows, row)
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	bot.Send(msg)
}
func (s *MessageSender) ShowStartButtons(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/question"),
			tgbotapi.NewKeyboardButton("/category"),
			tgbotapi.NewKeyboardButton("/exit"),
		),
	)
	bot.Send(msg)
}
