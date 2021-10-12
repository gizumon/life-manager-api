package actions

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"gizumon.com/life-manager-api/constants"
)

func TodoCmdAction(c echo.Context) (string, error) {
	message := "ごめんね、ToDoの機能はまだ開発中です🙇‍♂️"
	bot := c.Get(constants.LINE_HANDLER_KEY).(*linebot.Client)
	event := c.Get(constants.LINE_EVENT_KEY).(*linebot.Event)
	text := event.Message.(*linebot.TextMessage).Text

	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()

	fmt.Println("Message from user: ", text)

	if err != nil {
		return "failed", err
	}

	return "success", err
}
