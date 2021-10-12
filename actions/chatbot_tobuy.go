package actions

import (
	"github.com/labstack/echo/v4"
	"gizumon.com/life-manager-api/models"
	"gizumon.com/life-manager-api/requests"
	"gizumon.com/life-manager-api/responses"
)

type TobuyAction interface {
	Add(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error)
	GetOne(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error)
	GetList(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error)
	Delete(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error)
}

type tobuyAction struct {
	ToBuyRepository models.TobuyRepository
}

func NewTobuyAction(
	tobuyRepository models.TobuyRepository,
) TobuyAction {
	return &tobuyAction{
		ToBuyRepository: tobuyRepository,
	}
}

func (action *tobuyAction) Add(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error) {
	return &responses.SuccessResponse{
		Status:  "preparing",
		Results: []responses.ResultData{},
	}, nil
}

func (action *tobuyAction) GetOne(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error) {
	return &responses.SuccessResponse{
		Status:  "preparing",
		Results: []responses.ResultData{},
	}, nil
}

func (action *tobuyAction) GetList(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error) {
	return &responses.SuccessResponse{
		Status:  "preparing",
		Results: []responses.ResultData{},
	}, nil
}

func (action *tobuyAction) Delete(c echo.Context, args *requests.Args) (*responses.SuccessResponse, error) {
	return &responses.SuccessResponse{
		Status:  "preparing",
		Results: []responses.ResultData{},
	}, nil
}

// func TobuyCmdAction(c echo.Context) (string, error) {
// 	message := "ごめんね、ToBuyの機能はまだ開発中です🙇‍♂️"
// 	bot := c.Get(constants.LINE_HANDLER_KEY).(*linebot.Client)
// 	event := c.Get(constants.LINE_EVENT_KEY).(*linebot.Event)
// 	text := event.Message.(*linebot.TextMessage).Text

// 	args, err = actions.
// 	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()

// 	fmt.Println("Message from user: ", text)

// 	if err != nil {
// 		return "failed", err
// 	}

// 	return "success", err
// }
