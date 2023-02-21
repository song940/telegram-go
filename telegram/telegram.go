package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

type TelegramBotResponse struct {
	Ok          bool            `json:"ok"`
	Code        int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
	Result      json.RawMessage `json:"result"`
}

type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	UserName                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type Message struct {
	ChatId                string `json:"chat_id"`
	MessageThreadId       string `json:"message_thread_id,omitempty"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	Entities              string `json:"entities,omitempty"`
	DisableWebPageView    bool   `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool   `json:"disable_notification,omitempty"`
	ProtectContent        bool   `json:"protect_content,omitempty"`
	ReplyToMessageId      int    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithReply bool   `json:"allow_sending_with_reply,omitempty"`
}

type Update struct {
}

func NewBot(config *TelegramBotConfig) (bot *TelegramBot) {
	bot = &TelegramBot{
		config: config,
		client: http.DefaultClient,
	}
	return
}

func (bot *TelegramBot) Call(method string, params any) (result json.RawMessage, err error) {
	url := "https://api.telegram.org/bot" + bot.config.Token + method
	payload, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	req.Header.Add("content-type", "application/json")
	res, err := bot.client.Do(req)
	if err != nil {
		return
	}
	data, _ := io.ReadAll(res.Body)
	var out TelegramBotResponse
	err = json.Unmarshal(data, &out)
	result = out.Result
	if !out.Ok {
		return nil, fmt.Errorf("error: %d %s", out.Code, out.Description)
	}
	return
}

// GetMe
// https://core.telegram.org/bots/api#getme
func (bot *TelegramBot) GetMe() (user User, err error) {
	data, err := bot.Call("/getMe", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &user)
	return
}

// SendMessage
// https://core.telegram.org/bots/api#sendmessage
func (bot *TelegramBot) SendMessage(message Message) {
	data, err := bot.Call("/sendMessage", message)
	if err != nil {
		return
	}
	log.Println(data)
}

// GetUpdates
// https://core.telegram.org/bots/api#getting-updates
func (bot *TelegramBot) GetUpdates() {
	bot.Call("/getUpdates", nil)
}
