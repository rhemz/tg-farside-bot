package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
	"golang.org/x/net/context"
	"tg-farside-bot/config"
	"tg-farside-bot/handlers"
	"tg-farside-bot/logger"
	"time"
)

type (
	Server struct {
		echo *echo.Echo
		cfg  *config.Spec
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

func NewServer(cfg *config.Spec, logger zerolog.Logger) *Server {
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

	// routes
	e.GET("/health", handlers.HealthEndpoint)

	return &Server{
		echo: e,
		cfg:  cfg,
	}
}

func main() {
	cfg := config.GetConfig()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger.InitLogger(zerolog.DebugLevel, cfg.LogJSON) // TODO: cfg level

	server := NewServer(cfg, logger.Logger)
	server.Start()
}
