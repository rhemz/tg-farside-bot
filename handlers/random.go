package handlers

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"tg-farside-bot/config"
	"tg-farside-bot/logger"
)

func (h *Handler) RandomComicEndpoint(c echo.Context) error {
	chatId, err := strconv.ParseInt(c.Param("chatId"), 10, 64)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, Response{Message: "bad chat ID"}, JsonIndent)
	}

	randomComicYear := config.AllComics[rand.Intn(len(config.AllComics))]
	comicPath := strings.Join([]string{h.Config.S3BucketPath, randomComicYear.RandomImageFilename()}, "")

	result, err := h.S3.Client.GetObject(c.Request().Context(), &s3.GetObjectInput{
		Bucket: aws.String(h.Config.S3BucketName),
		Key:    aws.String(comicPath),
	})
	if err != nil {
		logger.L(c).Fatal().Err(err).Str("object_path", comicPath).Msg("Error getting object")
	}
	defer result.Body.Close()
	imageBytes, err := io.ReadAll(result.Body)

	photoData := tgbotapi.FileBytes{
		Name:  "thumbnail",
		Bytes: imageBytes,
	}
	photoMsg := tgbotapi.NewPhoto(chatId, photoData)
	photoMsg.ParseMode = "HTML"
	photoMsg.Caption = "test comic"
	resp, err := h.Api.Send(photoMsg)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, Response{
		Message: "ok",
		Data: struct {
			ChatId   int64 `json:"chat_id"`
			PostedAt int   `json:"posted_at"`
		}{
			ChatId:   resp.Chat.ID,
			PostedAt: resp.Date,
		},
	}, JsonIndent)
}
