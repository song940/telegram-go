package cli

import (
	"log"

	"github.com/song940/telegram-go/telegram"
)

func Run() {
	bot := telegram.NewBot(&telegram.TelegramBotConfig{
		Token: "830538223:AAG85jdDekD8RBD_hl4Fih0yeqnJBxn1bWcx",
	})
	me, err := bot.GetMe()
	log.Println(me, err)

	message := telegram.Message{
		Text: "Hello Telegram!",
	}
	bot.SendMessage(message)
}
