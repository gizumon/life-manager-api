package actions

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"gizumon.com/life-manager-api/constants"
)

func HelpCmdAction(c echo.Context) (string, error) {
	const helpText = "こんにちわ🐶下から知りたい項目を選んでね！"
	bot := c.Get(constants.LINE_HANDLER_KEY).(*linebot.Client)
	event := c.Get(constants.LINE_EVENT_KEY).(*linebot.Event)
	text := event.Message.(*linebot.TextMessage).Text

	_, err := bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(helpText).WithQuickReplies(
			linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton(constants.PAY_ICON_URL, linebot.NewMessageAction("Pay", "ヘルプ [pay]")),
				linebot.NewQuickReplyButton(constants.TODO_ICON_URL, linebot.NewMessageAction("ToDo", "ヘルプ [todo]")),
				linebot.NewQuickReplyButton(constants.TOBUY_ICON_URL, linebot.NewMessageAction("ToBuy", "ヘルプ [tobuy]")),
				linebot.NewQuickReplyButton(constants.HELP_ICON_URL, linebot.NewMessageAction("Help", "ヘルプ [tobuy]")),
			),
		),
	).Do()

	fmt.Println("Message from user: ", text)

	if err != nil {
		return "failed", err
	}

	return "success", err
}
