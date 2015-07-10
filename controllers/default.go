package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"github.com/chanxuehong/wechat/util"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"

	"net/http"
	// "net/url"

	. "maihaoguo/util"
)

type MainController struct {
	beego.Controller
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	//log.Println(err.Error())
}

// 文本消息的 Handler
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	mp.WriteRawResponse(w, r, resp) // 明文模式
	//mp.WriteAESResponse(w, r, resp) // 安全模式
}

func (c *MainController) Get() {

	signature := util.Sign("maihaoguo", c.GetString("timestamp"), c.GetString("nonce"))
	if signature == c.GetString("signature") {

		c.Ctx.WriteString(c.GetString("echostr"))

	} else {

	}

}

func (c *MainController) Post() {

	Log.Info("Post Request")
	aesKey, err := util.AESKeyDecode("6c0p5ooqHpxLm8tDQo2yFN0lPl8FiunHz6s3uwBHMyj") // 这里 encodedAESKey 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler)

	invalidRequesthandler := mp.InvalidRequestHandlerFunc(ErrorHandler)

	// 下面函数的几个参数设置成你自己的参数: oriId, token, appId
	mpServer := mp.NewDefaultServer("gh_d1ccedd04e4b", "maihaoguo", "wx8657df9e66c8277d", aesKey, messageServeMux)
	mpServerFrontend := mp.NewServerFrontend(mpServer, invalidRequesthandler, nil)

	// queryValues, err := url.ParseQuery(c.Ctx.Request.URL.RawQuery)
	// if err != nil {
	// 	Log.Info("错误暂不处理")
	// 	//frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
	// 	return
	// }

	// // if interceptor := frontend.interceptor; interceptor != nil && !interceptor.Intercept(w, r, queryValues) {
	// // 	return
	// // }

	// mp.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request, queryValues, mpServer, nil)

	mpServerFrontend.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}
