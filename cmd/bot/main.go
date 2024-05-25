package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/magickspell/go-bot/internal/app/commands"
	"github.com/magickspell/go-bot/internal/service/product"
)

// const token string = "TOKEN"

func main() {
	productService := product.NewService()

	godotenv.Load()

	var token string = os.Getenv("TOKEN")
	fmt.Println("token: " + token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	commander := commands.NewCommander(bot, productService)

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

			// switch update.Message.Command() {
			// case "help":
			// 	commander.Help(update.Message)
			// case "list":
			// 	commander.List(update.Message)
			// default:
			// 	commander.Default(update.Message)
			// }
			commander.HandleUpdate(update)
		}
	}
}
