package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"github.com/esap/wechat" // 微信SDK包
	"sort"
)

/*
wxef2acd768cace6db

5e61482aebf93740e1b887574e89ee03
*/
func main() {
	wechat.Debug = true

	cfg := &wechat.WxConfig{
		Token:          "yourToken",
		AppId:          "yourAppID",
		Secret:         "yourSecret",
		EncodingAESKey: "yourEncodingAesKey",
	}
	app := wechat.New(cfg)
	app.SendText("@all", "Hello,World!")

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	app.VerifyURL(w, r).NewText("客服消息1").Send().NewText("客服消息2").Send().NewText("查询OK").Reply()
	//})

	http.HandleFunc("/", checkout)

	http.ListenAndServe(":80", nil)
}



func checkout(response http.ResponseWriter, request *http.Request)  {
	//解析URL参数
	err := request.ParseForm()
	if err != nil {
		fmt.Println("URL解析失败！")
		return
	}
	// token
	var token string = "iwuqing"
	// 获取参数
	signature := request.FormValue("signature")
	timestamp := request.FormValue("timestamp")
	nonce := request.FormValue("nonce")
	echostr := request.FormValue("echostr")
	//将token、timestamp、nonce三个参数进行字典序排序
	var tempArray  = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature对比
	if sha1String == signature {
		_, err := response.Write([]byte(echostr))
		if err != nil {
			fmt.Println("响应失败。。。")
		}
	} else {
		fmt.Println("验证失败")
	}
}