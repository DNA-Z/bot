package commands

import (
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
