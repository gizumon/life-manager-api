package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gizumon.com/life-manager-api/responses"
)

func HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, responses.Health())
}
