package line_utils

import (
	"gizumon.com/life-manager-api/constants"
	"gizumon.com/life-manager-api/helpers"
)

func HasPayCmdKey(str string) bool {
	return helpers.IsInclude(str, constants.PAY_CMD_KEYS)
}

func HasTodoCmdKey(str string) bool {
	return helpers.IsInclude(str, constants.TODO_CMD_KEYS)
}

func HasTobuyCmdKey(str string) bool {
	return helpers.IsInclude(str, constants.TOBUY_CMD_KEYS)
}

func HasHelpCmdKey(str string) bool {
	return helpers.IsInclude(str, constants.HELP_CMD_KEYS)
}

func GetCmdKey(str string) string {
	return helpers.GetKeyFromValues(str, constants.CmdConversionMap)
}

func GetActionKey(str string) string {
	return helpers.GetKeyFromValues(str, constants.ActionConversionMap)
}

func GetCategoryKey(str string) string {
	return helpers.GetKeyFromValues(str, constants.CategoryConversionMap)
}

func GetPersonForKey(str string) string {
	return helpers.GetKeyFromValues(str, constants.CategoryConversionMap)
}

func GetPersonByKey(str string) string {
	return helpers.GetKeyFromValues(str, constants.CategoryConversionMap)
}

func GetPriceKey(str string) string {
	return helpers.GetNumberStrings(str)
}
