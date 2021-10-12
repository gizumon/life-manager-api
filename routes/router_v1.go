package routes

import (
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"

	// "gizumon.com/life-manager-api/constants"
	// "github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
	"gizumon.com/life-manager-api/actions"
	"gizumon.com/life-manager-api/handlers"
	"gizumon.com/life-manager-api/middlewares"
	"gizumon.com/life-manager-api/repositories"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.BodyDump(middlewares.RequestBodyDump))
	e.Use(middlewares.CustomContextMiddleware)
	chatbotHandler := handlers.NewChatbotHandler(
		actions.NewTobuyAction(
			repositories.NewTobuyRepository(),
		),
	)

	api := e.Group("/api/v1")
	// /api/v1/health
	api.GET("/health", handlers.HealthHandler)

	bot := api.Group("/line", middlewares.InitLine(), middlewares.InitFirebaseRDB())
	// /api/v1/line/notify
	// bot.POST("/notify", chatbotHandler.HandleMessage)
	// /api/v1/line/chatbot
	bot.POST("/chatbot", chatbotHandler.HandleMessage)

	return e
}
