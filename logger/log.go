package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	Logger zerolog.Logger
)

func InitLogger(level zerolog.Level, logJSON bool) {
	if logJSON {
		Logger = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(level)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		//zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	} else {
		// colorized stdout
		Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger().Level(level)
		zerolog.TimeFieldFormat = time.RFC3339
	}
}

func L(c echo.Context) *zerolog.Logger {
	return zerolog.Ctx(c.Request().Context())
}
