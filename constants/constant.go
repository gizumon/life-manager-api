package constants

import (
	"fmt"
	"os"
	"regexp"
)

// import (
// 	"regexp"
// 	"strings"
// )

const (
	// env keys
	ROOT_URL_KEY                            string = "ROOT_URL"
	API_PREFIX_KEY                          string = "API_PREFIX"
	CHANNEL_SECRET_KEY                      string = "CHANNEL_SECRET"
	CHANNEL_TOKEN_KEY                       string = "CHANNEL_TOKEN"
	FIREBASE_RDB_URL_KEY                    string = "FIREBASE_RDB_URL"
	GOOGLE_APPLICATION_CREDENTIALS_PATH_KEY string = "GOOGLE_APPLICATION_CREDENTIALS_PATH"

	// client key
	LINE_HANDLER_KEY        string = "LINE_HANDLER"
	LINE_EVENT_KEY          string = "LINE_EVENT"
	LINE_TARGET_USER_ID_KEY string = "LINE_TARGET_USER_ID"
	FIREBASE_RDB_CLIENT_KEY string = "FIREBASE_RDB_CLIENT"

	// icon url
	PAY_ICON_URL   string = ""
	TODO_ICON_URL  string = ""
	TOBUY_ICON_URL string = ""
	HELP_ICON_URL  string = "https://i.postimg.cc/DyMbPS0z/walking-Dog.png"
)

var (
	// cmd keys
	HELP_CMD_KEYS  = []string{"help", "Help", "Help", "へるぷ", "ヘルプ", "使い方", "❓"}
	PAY_CMD_KEYS   = []string{"pay", "PAY", "Pay", "支払", "払った", "¥", "💸"}
	TODO_CMD_KEYS  = []string{"todo", "TODO", "ToDo", "やる事", "タスク", "✓", "✅"}
	TOBUY_CMD_KEYS = []string{"tobuy", "TOBUY", "ToBuy", "Tobuy", "買い物", "かいもの", "🛒"}

	// action keys
	ARG_ACTION_ADD_KEYS    = []string{"add", "ADD", "追加", "+"}
	ARG_ACTION_LIST_KEYS   = []string{"list", "LIST", "show", "SHOW", "一覧", "ls"}
	ARG_ACTION_DELETE_KEYS = []string{"delete", "DELETE", "del", "DEL", "削除", "-"}

	// Separater
	ARG_SEPALATOR = regexp.MustCompile(`[¥s]+`)

	// directory
	BASE_PATH = getWorkingDirectory()
)

func getWorkingDirectory() string {
	d, err := os.Getwd()

	if err != nil {
		fmt.Errorf("Error in get working directory", err)
		return ""
	}

	return d
}
