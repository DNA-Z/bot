package commands

import (
	"log"

	"github.com/DNA-Z/bot/internal/services/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.ProductService
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.ProductService,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data: "+update.CallbackQuery.Data,
		)
		c.bot.Send(msg)
		return
	}

	if update.Message != nil {
		msg := update.Message

		switch msg.Command() {
		case "help":
			c.Help(msg)
		case "list":
			c.List(msg)
		case "get":
			c.Get(msg)
		default:
			c.Default(msg)
		}
	}
}
