package main

import (
	"strings"
)

func sayhi(str string, userID string) string {
	var result string
	switch userID {
	case "Uaa8f39e4ce4da3adb16e05cbcf52c919": //球球
		userWhom = "球球哥"
	case "U677c9c32e514ffa1ab1102e6e882315d": //天又
		userWhom = "天又哥"
	case "U0d1c8ed81dd9a9cfd402be5fb426e019": //姆斯
		userWhom = "james哥"
	case "Uf141820779edea05ac76f468f9634290": //子祐
		userWhom = "后世"
	case "U6cad7524aa07e15299355ee6173267af": //爸爸
		userWhom = "爸爸"
	case "Uf962a64b052b96faf22860b096e71643": //阿華
		userWhom = "阿华哥"
	case "U76fb3ac2fc172920e41873d8dbe615a1": //姐姐
		userWhom = "宁姊"
	case "U954ea584f53c0f4eb857f3b5ade713c3": //餅乾
		userWhom = "饼干"
	case "U6e8c5d2f59d2905b34b3bbccbba04733": //白白
		userWhom = "白白"
	case "U4739962db979ae81526e5e006f5f5174": //poyi
		userWhom = "poyi哥"
	case "U9f0778d04a28fcd15e3b1417bd5e308b": //紅茶
		userWhom = "大哥"
	}

	if strings.Contains(str, "G仔") {
		if strings.Contains(str, "早安") {
			result = userWhom + "早安啊"
		} else if strings.Contains(str, "午安") {
			result = userWhom + "午安啊"
		} else if strings.Contains(str, "晚安") {
			result = userWhom + "晚安啊"
		} else {
			result = userWhom + "怎么了@@?"
		}
	}

	return result
}
