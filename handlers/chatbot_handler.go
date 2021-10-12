package handlers

import (
	// "fmt"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	// "github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
	"gizumon.com/life-manager-api/actions"
	"gizumon.com/life-manager-api/constants"
	"gizumon.com/life-manager-api/helpers/line_utils"
	"gizumon.com/life-manager-api/models"
	"gizumon.com/life-manager-api/requests"
	"gizumon.com/life-manager-api/responses"
)

type ChatbotHandler interface {
	HandleMessage(c echo.Context) error
}

type chatbotHandler struct {
	TobuyAction actions.TobuyAction
}

// PayAction actions.PayAction
// TodoAction actions.TodoAction
// HelpAction actions.HelpAction
// ChatbotAction actions.ChatbotAction

func NewChatbotHandler(
	tobuyAction actions.TobuyAction,
) ChatbotHandler {
	return &chatbotHandler{
		TobuyAction: tobuyAction,
	}
}

// parse
func (h *chatbotHandler) HandleMessage(c echo.Context) error {
	responseData := []responses.ResultData{}
	bot := c.Get(constants.LINE_HANDLER_KEY).(*linebot.Client)
	events, err := bot.ParseRequest(c.Request())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.ErrorPlain("No Events", err.Error()))
	}

	for _, event := range events {
		data := responses.ResultData{}
		c.Set(constants.LINE_EVENT_KEY, event)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				res, err := handleTextMessage(message.Text, h, c)

				if err != nil {
					fmt.Print(err)
					return echo.NewHTTPError(http.StatusBadRequest, responses.ErrorPlain("Bad request", err.Error()))
				}

				data.Message = res.Results[0].Message
				data.Success = true
			}
			responseData = append(responseData, data)
		}
	}

	return c.JSON(http.StatusOK, responses.SuccessEvents(responseData))
}

func handleTextMessage(text string, h *chatbotHandler, c echo.Context) (*responses.SuccessResponse, error) {
	var res *responses.SuccessResponse = &responses.SuccessResponse{}
	var err error = nil
	args := textSeparateToArgs(text)

	switch args.Cmd {
	case "tobuy":
		res, err = handleTobuyAction(args, h, c)
	default:
		err = errors.New("unknown cmd key")
	}

	return res, err
}

func handleTobuyAction(args *requests.Args, h *chatbotHandler, c echo.Context) (*responses.SuccessResponse, error) {
	switch args.Action {
	case models.Add:
		return h.TobuyAction.Add(c, args)
	case models.List:
		return h.TobuyAction.GetList(c, args)
	default:
		return nil, errors.New("unknown action")
	}
}

func textSeparateToArgs(text string) *requests.Args {
	switch {
	// case line_utils.HasPayCmdKey(text):
	// 	return requests.SeparateTextToPayArgs(text)
	case line_utils.HasTobuyCmdKey(text):
		return requests.SeparateTextToTobuyArgs(text)
	// case line_utils.HasTodoCmdKey(text):
	// 	return requests.SeparateTextToTodoArgs(text)
	case line_utils.HasHelpCmdKey(text):
		return &requests.Args{
			CommonArgs: requests.CommonArgs{Cmd: "@help", Action: ""},
		}
	default:
		return &requests.Args{}
	}
}
