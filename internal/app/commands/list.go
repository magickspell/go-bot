package commands

import (
	"encoding/json"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	products := c.productService.List()
	var outMsg string = "All products: \n\n"
	for n, prod := range products {
		outMsg += strconv.Itoa(n) + ". " + prod.Title + "\n"
	}

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		outMsg,
	)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)
	c.bot.Send(msg)
}

func init() {
	registredCommands["list"] = (*Commander).List
}
