package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"math/rand"
	"tg-farside-bot/config"
	"tg-farside-bot/logger"
	"tg-farside-bot/util"
	"time"
)

func main() {
	cfg := config.GetConfig()

	logger.InitLogger(zerolog.DebugLevel, cfg.LogJSON) // TODO: cfg level

	tgApi, err := tgbotapi.NewBotAPI(config.GetConfig().TelegramBotToken)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("Error initializing telegram bot API client")
	}

	s3client, err := util.NewDOSpacesS3Client(cfg.S3Endpoint, cfg.S3KeyId, cfg.S3SecretKey)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("Error initializing S3 client")
	}

	// seed PRNG
	rand.Seed(time.Now().UnixNano())

	// glhf gogo
	s := NewServer(cfg, logger.Logger, tgApi, s3client)
	s.Start()
}
