package handlers

import (
	// "fmt"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	// "github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
	"gizumon.com/life-manager-api/actions"
	"gizumon.com/life-manager-api/constants"
	"gizumon.com/life-manager-api/responses"
)

func ChatbotHandler(c echo.Context) error {
	responseData := []responses.ResultData{}
	bot := c.Get(constants.LINE_HANDLER_KEY).(*linebot.Client)
	events, err := bot.ParseRequest(c.Request())

	if err != nil {
		return err
	}

	for _, event := range events {
		data := responses.ResultData{}
		c.Set(constants.LINE_EVENT_KEY, event)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				res, err := handleWithMessage(message.Text, c)
				data.Message = res
				data.Success = true
				if err != nil {
					fmt.Print(err)
					data.Success = false
				}
			}
			responseData = append(responseData, data)
		}
	}

	return c.JSON(http.StatusOK, responses.SuccessEvents(responseData))
}

func handleWithMessage(text string, c echo.Context) (string, error) {
	if hasPayCmdKey(text) {
		return actions.PayCmdAction(c)
	}
	if hasTodoCmdKey(text) {
		return actions.HelpCmdAction(c)
	}
	if hasTobuyCmdKey(text) {
		return actions.HelpCmdAction(c)
	}
	if hasHelpCmdKey(text) {
		return actions.HelpCmdAction(c)
	}
	return actions.HelpCmdAction(c)
}

func hasPayCmdKey(str string) bool {
	for _, key := range constants.PAY_CMD_KEYS {
		if strings.Contains(str, key) {
			return true
		}
	}
	return false
}

func hasTodoCmdKey(str string) bool {
	for _, key := range constants.TODO_CMD_KEYS {
		if strings.Contains(str, key) {
			return true
		}
	}
	return false
}

func hasTobuyCmdKey(str string) bool {
	for _, key := range constants.TOBUY_CMD_KEYS {
		if strings.Contains(str, key) {
			return true
		}
	}
	return false
}

func hasHelpCmdKey(str string) bool {
	for _, key := range constants.HELP_CMD_KEYS {
		if strings.Contains(str, key) {
			return true
		}
	}
	return false
}
