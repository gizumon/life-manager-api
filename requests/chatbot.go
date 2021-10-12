package requests

import (
	"strings"

	"gizumon.com/life-manager-api/constants"
	"gizumon.com/life-manager-api/helpers/line_utils"
)

type LineMessage struct {
	Text string
}

type Args struct {
	CommonArgs
	PayArgs
	TobuyArgs
	TodoArgs
}

// TODO: Update actual parameter
// [Cmd] [Action] [Item] [other]
// ex) pay add 1000 sugar @person
// export interface IPay {
// 	id: string;
// 	price: number;
// 	payedFor: string[];
// 	payedCategory: string;
// 	payedBy: string;
// 	memo: string;
// 	timestamp: number;
// }

// @pay (add) 1000 what me/you
type CommonArgs struct {
	Cmd    string // tobuy | use CmdConversionMap
	Action string // add, list, delete | use ActionConversionMap
}

type PayArgs struct {
	// ID       string `json:"id"`
	Price       string `json:"price" validate:"int"` // 1000
	PayItem     string `json:"memo"`                 // [any] | no conversionMap
	For         string `json:"payedFor"`             // all | partner
	PayCategory string `json:"payedCategory"`        // [any of categories] | CategoryConversionMap
	PayBy       string `json:"payedBy"`              // [any of categories] | CategoryConversionMap
}

// [Cmd] [Action] [Item] [other]
// ex) @tobuy add sugar other
type TobuyArgs struct {
	BuyItem     string `json:"item"`          // [any] | no conversionMap
	BuyCategory string `json:"tobuyCategory"` // [any of categories] | CategoryConversionMap
	// By       string //
	// Due      string //
}

// [Cmd] [Action] [Item] [other]
// ex) @tobuy add sugar other
type TodoArgs struct {
	DoItem string // [any] | no conversionMap
	DoBy   string //
	Due    string //
}

var DefaultPayArgs = PayArgs{
	Price:       "0",
	PayItem:     "",
	For:         "all",
	PayCategory: "none",
	PayBy:       "me",
}

var DefaultTobuyArgs = TobuyArgs{
	BuyItem:     "default",
	BuyCategory: "none",
}

var DefaultTodoArgs = TodoArgs{
	DoItem: "default",
	DoBy:   "me",
}

func GetCmdKey(text string) string {
	return line_utils.GetCmdKey(text)
}

func SeparateTextToPayArgs(text string) *Args {
	var args Args
	text = strings.Trim(text, " ")
	strList := constants.ARG_SEPALATOR.Split(text, -1)
	for i := 0; i < len(strList); i++ {
		cmdKey := line_utils.GetCmdKey(strList[i])
		actionKey := line_utils.GetActionKey(strList[i])
		categoryKey := line_utils.GetCategoryKey(strList[i])
		personForKey := line_utils.GetPersonForKey(strList[i])
		personByKey := line_utils.GetPersonByKey(strList[i])
		priceKey := line_utils.GetPriceKey(strList[i])

		switch {
		case cmdKey != "":
			args.Cmd = cmdKey
		case actionKey != "":
			args.Action = actionKey
		case categoryKey != "":
			args.BuyCategory = categoryKey
		case priceKey != "":
			args.Price = priceKey
		case personForKey != "":
			args.For = personForKey
		case personByKey != "":
			args.PayBy = personByKey
		default:
			args.PayItem = strList[i]
		}
	}
	return &args
}

func SeparateTextToTobuyArgs(text string) *Args {
	var args Args
	text = strings.Trim(text, " ")
	strList := constants.ARG_SEPALATOR.Split(text, -1)
	for i := 0; i < len(strList); i++ {
		cmdKey := line_utils.GetCmdKey(strList[i])
		actionKey := line_utils.GetActionKey(strList[i])
		categoryKey := line_utils.GetCategoryKey(strList[i])
		// personForKey := line_utils.GetPersonForKey(strList[i])
		// personByKey := line_utils.GetPersonByKey(strList[i])
		switch {
		case cmdKey != "":
			args.Cmd = cmdKey
		case actionKey != "":
			args.Action = actionKey
		case categoryKey != "":
			args.BuyCategory = categoryKey
		default:
			args.BuyItem = strList[i]
		}
	}
	return &args
}

func SeparateTextToTodoArgs(text string) *Args {
	var args Args
	text = strings.Trim(text, " ")
	strList := constants.ARG_SEPALATOR.Split(text, -1)
	for i := 0; i < len(strList); i++ {
		cmdKey := line_utils.GetCmdKey(strList[i])
		actionKey := line_utils.GetActionKey(strList[i])
		// personForKey := line_utils.GetPersonForKey(strList[i])
		// personByKey := line_utils.GetPersonByKey(strList[i])
		switch {
		case cmdKey != "":
			args.Cmd = cmdKey
		case actionKey != "":
			args.Action = actionKey
		default:
			args.DoItem = strList[i]
		}
	}
	return &args
}
