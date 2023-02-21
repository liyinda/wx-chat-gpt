package main

import (
	"net/http"
	"github.com/esap/wechat" // 微信SDK包
)

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.VerifyURL(w, r).NewText("客服消息1").Send().NewText("客服消息2").Send().NewText("查询OK").Reply()
	})

	http.ListenAndServe(":9090", nil)
}