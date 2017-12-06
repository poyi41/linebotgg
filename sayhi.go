package main

import (
	"strings"
)

func sayhi(str string, userID string) string {
	str = strings.ToUpper(str)
	var result string
	var userWhom string
	userWhom = getName(userID)

	if strings.Contains(str, "G仔") {
		result = userWhom + "怎么了@@?"

		if strings.Contains(str, "午安") {
			result = userWhom + "午安丫"
		} else if strings.Contains(str, "早安") {
			result = userWhom + "早安丫"
		} else if strings.Contains(str, "晚安") {
			result = userWhom + "晚安丫"
		} else if strings.Contains(str, "在幹嘛") {
			result = "趕稿去= 3="
		} else if strings.Contains(str, "在幹嘛") {
			result = userWhom + "好久不見，哇咔咔"
		}
	} else if strings.Contains(str, "中文豬") {
		result = userWhom + "别这样= 3="
	}

	return result
}
