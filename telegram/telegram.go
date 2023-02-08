package telegram

import (
	"fmt"
	"io"
	"net/http"
)

type TelegramBotConfig struct {
	API   string `json:"api"`
	Token string `json:"token"`
}

type TelegramBot struct {
	config *TelegramBotConfig
	client *http.Client
}

func NewBot(config *TelegramBotConfig) (bot *TelegramBot) {
	bot = &TelegramBot{
		config: config,
		client: http.DefaultClient,
	}
	return
}

func (bot *TelegramBot) Call(method string, body io.Reader) (res interface{}, err error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.config.Token, method)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	res, err = bot.client.Do(req)
	return
}

func (bot *TelegramBot) GetMe() (res interface{}, err error) {
	return bot.Call("/getMe", nil)
}

func (bot *TelegramBot) GetUpdates() (res interface{}, err error) {
	return
}
