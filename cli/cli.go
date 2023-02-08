package cli

import (
	"log"

	"github.com/song940/telegram/telegram"
)

func Run() {
	bot := telegram.NewBot(&telegram.TelegramBotConfig{
		Token: "",
	})
	me, err := bot.GetMe()
	log.Println(me, err)
}
