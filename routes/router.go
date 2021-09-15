package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"gizumon.com/life-manager-api/handlers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(mw.Logger())

	api := e.Group("/api")
	// api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/health", handlers.HealthHandler)
	api.POST("/webhook", handlers.WebhookHandler)

	return e
}
