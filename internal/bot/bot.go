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

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	repo := repository.NewInMemoryQuestionRepository()
	service := services.NewQuestionService(repo)
	handler := NewBotHandler(bot, service)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	handler.HandleUpdates(updates)
}
