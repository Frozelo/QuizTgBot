package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *BotHandler) handleNormalState(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		h.sendWelcomeMessage(message.Chat.ID)
	case "question":
		h.startTest(message.Chat.ID)
	case "test":
		h.startTest(message.Chat.ID)
	case "category":
		h.promptForCategory(message.Chat.ID)
	case "exit":
		h.endTest(message.Chat.ID)
	default:
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Используйте /start, /question, /category или /exit."))
	}
}

func (h *BotHandler) handleInTestState(message *tgbotapi.Message) {
	if message.Command() == "exit" {
		h.endTest(message.Chat.ID)
	} else {
		h.handleAnswer(message)
	}
}

func (h *BotHandler) handleAwaitingCategoryState(message *tgbotapi.Message) {
	h.category = strings.TrimSpace(message.Text)
	h.state = InTestState
	h.askRandomQuestionByCategory(message.Chat.ID)
}
