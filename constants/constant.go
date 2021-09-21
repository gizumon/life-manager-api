package constants

// import (
// 	"regexp"
// 	"strings"
// )

const (
	// env keys
	ROOT_URL_KEY       string = "ROOT_URL"
	API_PREFIX_KEY     string = "API_PREFIX"
	CHANNEL_SECRET_KEY string = "CHANNEL_SECRET"
	CHANNEL_TOKEN_KEY  string = "CHANNEL_TOKEN"

	// client key
	LINE_HANDLER_KEY string = "LINE_HANDLER"
	LINE_EVENT_KEY   string = "LINE_EVENT"

	// icon url
	// PAY_ICON_URL   string = "http://www.entypo.com/images/squared-minus.svg"
	// TODO_ICON_URL  string = "http://www.entypo.com/images/squared-plus.svg"
	// TOBUY_ICON_URL string = "http://www.entypo.com/images/shopping-cart.svg"
	// HELP_ICON_URL  string = "http://www.entypo.com/images/help-with-circle.svg"
	PAY_ICON_URL   string = ""
	TODO_ICON_URL  string = ""
	TOBUY_ICON_URL string = ""
	HELP_ICON_URL  string = "https://i.postimg.cc/DyMbPS0z/walking-Dog.png"
)

var (
	HELP_CMD_KEYS  = []string{"help", "Help", "Help", "へるぷ", "ヘルプ", "使い方", "❓"}
	PAY_CMD_KEYS   = []string{"pay", "PAY", "Pay", "支払", "払った", "¥", "💸"}
	TODO_CMD_KEYS  = []string{"todo", "TODO", "ToDo", "やる事", "タスク", "✓", "✅"}
	TOBUY_CMD_KEYS = []string{"tobuy", "TOBUY", "ToBuy", "Tobuy", "買い物", "かいもの", "🛒"}
)
