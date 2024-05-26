package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/magickspell/go-bot/internal/service/product"
)

var registredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot: bot,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	// mutex := sync.Mutex{} // mutex - это способ организации синхронной работы с данными между горутинами
	// mutex.Lock()
	// defer mutex.Unlock()
	// defer deferFunc() // функция выполнится даже когда все упало (паника например panic())
	defer func() { // но не нужно обрабатывать исключения через эту штуку!
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic")
		}
	}()

	command, ok := registredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}
}
