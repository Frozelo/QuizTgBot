package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type MessageSender struct{}

func NewMessageSender() *MessageSender {
	return &MessageSender{}
}

func (s *MessageSender) SendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Привет! Я бот для подготовки к собеседованиям по Go. Используйте команду /question для получения вопроса. Используйте команду /category для выбора категории вопросов.")
	bot.Send(msg)
	s.ShowStartButtons(bot, chatID)
}

func (s *MessageSender) SendQuestionMessage(bot *tgbotapi.BotAPI, chatID int64, question string) {
	msg := tgbotapi.NewMessage(chatID, question)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ответ 1", "Ответ 1"),
			tgbotapi.NewInlineKeyboardButtonData("Ответ 2", "Ответ 2"),
			tgbotapi.NewInlineKeyboardButtonData("Ответ 3", "Ответ 3"),
		),
	)
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
