package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) TelegramEventEndpoint(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, Response{Message: "ok"}, JsonIndent)
}
