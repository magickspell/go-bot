package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// const token string = "TOKEN"

func main() {
	godotenv.Load()

	var token string = os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u) // очередь апдейтов

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Println("")
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote:"+update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID // делаем реплай на введеное сообщение

			bot.Send(msg)
		}
	}
}
