package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) TelegramEventEndpoint(c echo.Context) error {
	var update tgbotapi.Update

	err := c.Bind(&update)
	if err != nil {
		// even if error binding, always return 200 to not trigger telegram infinite redelivery
		return c.JSONPretty(http.StatusOK, Response{Message: "error reading JSON"}, JsonIndent)
	}

	return c.JSONPretty(http.StatusOK, Response{Message: "ok"}, JsonIndent)

}
