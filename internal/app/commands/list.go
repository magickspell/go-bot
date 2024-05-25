package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	products := c.productService.List()
	var outMsg string = "All products: \n\n"
	for n, prod := range products {
		outMsg += strconv.Itoa(n) + ". " + prod.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		outMsg,
	)
	c.bot.Send(msg)
}

func init() {
	registredCommands["list"] = (*Commander).List
}
