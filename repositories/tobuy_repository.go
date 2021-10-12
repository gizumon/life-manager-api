package repositories

import (
	"strings"

	"gizumon.com/life-manager-api/constants"
	"gizumon.com/life-manager-api/helpers/line_utils"
	"gizumon.com/life-manager-api/models"
)

type tobuyRepository struct{}

func NewTobuyRepository() models.TobuyRepository {
	return &tobuyRepository{}
}

var DefaultArgs = models.TobuyArgs{
	Cmd:      "@buy",
	Action:   "add",
	Item:     "default",
	Category: "none",
}

// 1. How many args in texts
func (re *tobuyRepository) SeparateTextToArgs(text string) (*models.TobuyArgs, error) {
	// @buy add Shimeji none
	text = strings.Trim(text, " ")
	strList := constants.ARG_SEPALATOR.Split(text, -1)
	args := DefaultArgs
	// flagMap := map[string]bool{
	// 	"Cmd":      false,
	// 	"Action":   false,
	// 	"Item":     false,
	// 	"Category": false,
	// }

	for i := 0; i < len(strList); i++ {
		cmdKey := line_utils.GetCmdKey(strList[i])
		actionKey := line_utils.GetActionKey(strList[i])
		categoryKey := line_utils.GetCategoryKey(strList[i])
		switch {
		case cmdKey != "":
			args.Cmd = cmdKey
			// flagMap["Cmd"] = true
		case actionKey != "":
			args.Action = actionKey
			// flagMap["Action"] = true
		case categoryKey != "":
			args.Category = categoryKey
			// flagMap["Category"] = true
		default:
			args.Item = strList[i]
			// flagMap["Item"] = true
		}
	}
	return &args, nil
}

func (re *tobuyRepository) Add(args *models.TobuyArgs) (string, error) {
	return "success", nil
}

func (re *tobuyRepository) GetOne(tobuyID string) (*models.Tobuy, error) {
	return &models.Tobuy{}, nil
}

func (re *tobuyRepository) GetList() (*[]models.Tobuy, error) {
	return &[]models.Tobuy{}, nil
}

func (re *tobuyRepository) Delete(tobuyID string) (string, error) {
	return "success", nil
}
