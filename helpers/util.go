package helpers

import (
	"regexp"
	"strings"
)

func IsInclude(val string, arr []string) bool {
	for _, key := range arr {
		if strings.Contains(val, key) {
			return true
		}
	}
	return false
}

func GetKeyFromValues(searchVal string, obj map[string][]string) string {
	for key, values := range obj {
		if IsInclude(searchVal, values) {
			return key
		}
	}
	return ""
}

func GetNumberStrings(str string) string {
	return regexp.MustCompile(`^¥d+$`).FindString(str)
}
