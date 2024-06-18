package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *BotHandler) showCategoryButtons(chatID int64) {
	categories, err := h.questionService.GetCategories()
	if err != nil {
		h.bot.Send(tgbotapi.NewMessage(chatID, "Не удалось получить категории вопросов."))
		return
	}

	var buttons []tgbotapi.KeyboardButton
	for _, category := range categories {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(category))
	}

	replyMarkup := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(buttons...),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите категорию:")
	msg.ReplyMarkup = replyMarkup

	h.bot.Send(msg)
}

func (h *BotHandler) showStartButtons(chatID int64) {
	buttons := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("/question"),
		tgbotapi.NewKeyboardButton("/test"),
		tgbotapi.NewKeyboardButton("/category"),
	}

	replyMarkup := tgbotapi.NewReplyKeyboard(buttons)
	msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
	msg.ReplyMarkup = replyMarkup

	h.bot.Send(msg)
}
