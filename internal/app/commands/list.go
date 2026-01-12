package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	product := c.productService.List()
	for i, p := range product {
		rang := i + 1
		outputMsgText += strconv.Itoa(rang) + ". " + p.Title
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}
