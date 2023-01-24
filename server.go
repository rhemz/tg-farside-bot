package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
	"golang.org/x/net/context"
	"tg-farside-bot/config"
	"tg-farside-bot/handlers"
	"tg-farside-bot/util"
	"time"
)

type (
	Server struct {
		echo        *echo.Echo
		cfg         *config.Spec
		telegramApi *tgbotapi.BotAPI
		s3Client    *util.DOSpacesS3Client
	}
)

func (s *Server) Start() {
	//go func() {
	//	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%d", s.cfg.ListenPort)))
	//}()
	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%d", s.cfg.ListenPort)))
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal(errors.Wrap(err, "failed to shutdown server"))
	}
}

func NewServer(cfg *config.Spec, logger zerolog.Logger, api *tgbotapi.BotAPI, s3 *util.DOSpacesS3Client) *Server {
	echoLogger := lecho.From(logger)

	e := echo.New()
	e.Debug = cfg.EchoDebug
	e.Logger = echoLogger
	e.HideBanner = cfg.LogJSON

	e.Use(middleware.RequestID())
	e.Use(lecho.Middleware(lecho.Config{
		Logger:       echoLogger,
		RequestIDKey: "request_id",
		NestKey:      "request",
	}))

	h := &handlers.Handler{
		Api:    api,
		Config: cfg,
		S3:     s3,
	}

	// routes
	e.GET("/health", h.HealthEndpoint)
	e.POST("/v1/event", h.TelegramEventEndpoint)
	e.POST("/v1/random-comic/:chatId", h.RandomComicEndpoint)

	return &Server{
		echo:        e,
		cfg:         cfg,
		telegramApi: api,
		s3Client:    s3,
	}
}
