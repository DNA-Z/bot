package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Printf("Wrong args", args)
	}

	product, err := c.productService.Get(idx)

	if err != nil {
		log.Println("Fail to get product with idx %d: %v", idx, args)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	c.bot.Send(msg)
}
