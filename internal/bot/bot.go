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

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	questionService := services.NewQuestionService(repo)
	messageSender := NewMessageSender()
	stateHandler := NewStateHandler(questionService, messageSender)

	for update := range updates {
		if update.Message != nil {
			stateHandler.HandleState(bot, update.Message)
		} else if update.CallbackQuery != nil {
			stateHandler.HandleCallback(bot, update.CallbackQuery)
		}
	}
}
