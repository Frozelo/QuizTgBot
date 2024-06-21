package bot

import (
	"log"
	"os"
	"quiz-bot/internal/domain/services"
	"quiz-bot/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	repo := repository.NewInMemoryQuestionRepository()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	questionService := services.NewQuestionService(repo)
	messageSender := NewMessageSender()
	bot := NewBot(botAPI, questionService, messageSender)

	// Настройка обработчиков сообщений и callback'ов
	botAPI.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := botAPI.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			bot.HandleMessage(update.Message)
		}
		if update.CallbackQuery != nil {
			bot.HandleCallback(update.CallbackQuery)
		}
	}
}
