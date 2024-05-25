package commands

import (
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
	command, ok := registredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}
}
