package controllers

import (
	"github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/util"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	log.SetLogger("console", "")
	signature := util.Sign(c.GetString("token"), c.GetString("timestamp"), c.GetString("nonce"))

	if signature == c.GetString("signature") {

		c.Ctx.WriteString(c.GetString("echostr"))
		log.Info(signature)

	} else {

	}

}
