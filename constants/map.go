package constants

var (
	CmdConversionMap = map[string][]string{
		"pay":   PAY_CMD_KEYS,
		"tobuy": TOBUY_CMD_KEYS,
		"todo":  TODO_CMD_KEYS,
	}
	ActionConversionMap = map[string][]string{
		"add":    ARG_ACTION_ADD_KEYS,
		"list":   ARG_ACTION_LIST_KEYS,
		"delete": ARG_ACTION_DELETE_KEYS,
	}
	// TODO: Use firebase data
	CategoryConversionMap = map[string][]string{
		"none":          {"none", "その他", "そのた", "other", "❔", "?"},
		"super":         {"super", "スーパー", "すーぱー"},
		"drugstore":     {"drugstore", "薬局", "薬", "💊"},
		"oneDollarShop": {"oneDollar", "100", "百均", "１００", "ひゃっきん"},
	}
	// to add group partner name
	PersonForConversionMap = map[string][]string{
		"all":     {"all", "ALL", "全員", "皆", "みんな"},
		"partner": {"partner", "相手", "あいて", ""},
	}
	// get from firebase
	PersonByConversionMap = map[string][]string{
		"me":      {"me", "自分", "相手"},
		"partner": {"partner", "相手", "あいて", ""},
	}
)
