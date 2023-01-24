package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-farside-bot/config"
	"tg-farside-bot/util"
)

type (
	Handler struct {
		Api    *tgbotapi.BotAPI
		Config *config.Spec
		S3     *util.DOSpacesS3Client
	}

	Response struct {
		Message string `json:"message,omitempty"`
		Data    any    `json:"data,omitempty"`
	}
)
