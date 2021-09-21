package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gizumon.com/life-manager-api/responses"
)

func NotifyHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, responses.Health())
}
