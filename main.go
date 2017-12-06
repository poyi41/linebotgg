// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				returnMsg(message.Text, event.Source.UserID, event.ReplyToken, err)
			}
		}
	}
}

func returnMsg(str string, userID string, rpyToken string, errStr error) {
	var result string
	var userWhom string
	switch userID {
	case "Uaa8f39e4ce4da3adb16e05cbcf52c919": //球球
		userWhom = "球球"
	case "U677c9c32e514ffa1ab1102e6e882315d": //天又
		userWhom = "天又"
	case "U0d1c8ed81dd9a9cfd402be5fb426e019": //姆斯
		userWhom = "姆斯"
	case "Uf141820779edea05ac76f468f9634290": //子祐
		userWhom = "子祐"
	case "U6cad7524aa07e15299355ee6173267af": //爸爸
		userWhom = "爸爸"
	case "Uf962a64b052b96faf22860b096e71643": //阿華
		userWhom = "阿華"
	case "U76fb3ac2fc172920e41873d8dbe615a1": //姐姐
		userWhom = "姐姐"
	case "U954ea584f53c0f4eb857f3b5ade713c3": //餅乾
		userWhom = "餅乾"
	case "U6e8c5d2f59d2905b34b3bbccbba04733": //白白
		userWhom = "白白"
	case "U4739962db979ae81526e5e006f5f5174": //poyi
		userWhom = "poyi"
		result = poyiReturn(str)
	case "U9f0778d04a28fcd15e3b1417bd5e308b": //紅茶
		userWhom = "紅茶"
	default:
		userWhom = userID
	}
	/*
		if strings.Contains(str, "G仔") {
			result = "@@"
		}else if strings.Contains(str, "G在幹嘛") {
			result = "趕稿去= 3="
		}

		switch str {
		case "中文豬":
			result = "= 3="
		case "早安":
			result = "大家早安"
		case "@@":
			result = "哇咔咔"
		default:
			log.Print(userID+": "+str)
		}
	*/
	log.Print(userWhom + ": " + str)

	if len(result) != 0 {
		if _, errStr = bot.ReplyMessage(rpyToken, linebot.NewTextMessage(result)).Do(); errStr != nil {
			log.Print(errStr)
		}
	}
}
