package main

import (
	"strings"
)

func sayhi(str string, userID string) string {
	str = strings.ToUpper(str)
	var result string
	var userWhom string
	userWhom = getName(userID)

	if strings.Contains(str, "中文豬") {
		result = userWhom + "别这样= 3="
	} else if strings.Contains(str, "@@") {
		result = "＠＠"
	} else if strings.Contains(str, "=3=") {
		result = "= 3="
	} else if strings.Contains(str, "嗨洞") {
		result = "((o))"
	} else if strings.Contains(str, "餅乾") {
		result = "我宁愿摩擦饼干"
	} else if strings.Contains(str, "慶記") {
		result = "等我去台湾后一定要带我吃一次庆記"
	} else if strings.Contains(str, "朱國強") {
		result = userWhom + "很故意= 3="
	} else if strings.Contains(str, "爐石") && strings.Contains(str, "任務") {
		result = "所以我说有人要帮解炉石任务吗？如果没有我晚点再问一次"
	} else if strings.Contains(str, "生日快樂") {
		result = "生日快乐啊。我的心是很好吃。我的朋友都在想什么。我的朋友们一起分享吧。我的心都碎了吗。我们要是你的朋友都说不出来有关部门申请加入会员抽大奖？我们的心理状态？我们要是我的朋友也能像这次我们的朋友也都非常喜欢她吗？我们要是有人要一起努力过后才知道是什么样子。我的心都碎了。我的朋友都在看电视了。我的心是一样的"
	} else if strings.Contains(str, "G仔") {

		switch userWhom {
		case "大哥":
			result = "大哥送我贴图=3="
		case "白白":
			result = "怎样@@?"
		case "james哥":
			result = strings.Replace(str, "G仔", userWhom, -1)
		default:
			result = userWhom + "怎么了@@?"
		}

		if strings.Contains(str, "沒事") && userWhom == "天又哥" {
			result = "没事就好"
		}

		if strings.Contains(str, "沒怎樣") && userWhom == "白白" {
			result = "没怎样就好"
		}

		if strings.Contains(str, "午安") {
			result = userWhom + "午安丫"
		} else if strings.Contains(str, "早安") {
			result = userWhom + "早安丫"
		} else if strings.Contains(str, "晚安") {
			result = userWhom + "晚安丫"
		} else if strings.Contains(str, "在幹嘛") {
			result = "趕稿去= 3="
		} else if strings.Contains(str, "好久不見") {
			result = userWhom + "好久不見，哇咔咔"
		} else if strings.Contains(str, "吃飯沒") {
			result = "吃了丫，" + userWhom + "你呢@@？"
		}
	}
	if str == "為所欲為" {
		result = "為所欲為"
	}

	if str == "是不是" {
		result = "是不是！！"
	}
	return result
}

func getName(userID string) string {
	var userWhom string
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
	case "U7ea7e8a2300bfbaaf58108af1786cc89": //柳丁
		userWhom = "柳丁"
	}
	return userWhom
}
