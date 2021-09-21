package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	// "github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
	"gizumon.com/life-manager-api/constants"
)

func LineAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			client, err := linebot.New(
				os.Getenv(constants.CHANNEL_SECRET_KEY),
				os.Getenv(constants.CHANNEL_TOKEN_KEY),
			)
			if err != nil {
				return echo.ErrInternalServerError
			}

			c.Set(constants.LINE_HANDLER_KEY, client)
			return next(c)
		}
	}
}
