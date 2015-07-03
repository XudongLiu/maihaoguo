package main

import (
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
	_ "maihaoguo/routers"
	. "maihaoguo/util"

	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
)

var AccessTokenServer = mp.NewDefaultAccessTokenServer("wx8657df9e66c8277d", "2b94ba328f84216c4eb2bfa0a6cbce82", nil) // 一個應用只能有一個實例
var mpClient = mp.NewClient(AccessTokenServer, nil)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}

// 文本消息的 Handler
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	//mp.WriteRawResponse(w, r, resp) // 明文模式
	mp.WriteAESResponse(w, r, resp) // 安全模式
}

func main() {
	beego.Run()
	var mn menu.Menu
	mn.Buttons = make([]menu.Button, 3)
	mn.Buttons[0].SetAsClickButton("我买果", "V1001_TODAY_MUSIC")
	mn.Buttons[1].SetAsViewButton("我卖果", "http://v.qq.com/")

	var subButtons = make([]menu.Button, 2)
	subButtons[0].SetAsViewButton("京郊水果", "http://www.soso.com/")
	subButtons[1].SetAsClickButton("进口精品", "V1001_GOOD")

	mn.Buttons[2].SetAsSubMenuButton("加入我们", subButtons)

	clt := menu.Client{Client: mpClient}
	if err := clt.CreateMenu(mn); err != nil {
		Log.Info(err.Error())
		return
	}
	Log.Info("菜单创建完成")

	aesKey, err := util.AESKeyDecode("6c0p5ooqHpxLm8tDQo2yFN0lPl8FiunHz6s3uwBHMyj") // 这里 encodedAESKey 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: oriId, token, appId
	mpServer := mp.NewDefaultServer("gh_d1ccedd04e4b", "maihaoguo", "wx8657df9e66c8277d", aesKey, messageServeMux)

	mpServerFrontend := mp.NewServerFrontend(mpServer, mp.ErrorHandlerFunc(ErrorHandler), nil)

	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/wechat
	// 那么可以这么注册 http.Handler
	http.Handle("/wechat", mpServerFrontend)
	//http.ListenAndServe(":80", nil)

}
