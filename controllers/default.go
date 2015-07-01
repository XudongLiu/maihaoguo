package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/chanxuehong/wechat/util"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
	signature := util.Sign("maihaoguo", c.GetString("timestamp"), c.GetString("nonce"))
	log.Info(c.GetString("signature"))
	if signature == c.GetString("signature") {

		c.Ctx.WriteString(c.GetString("echostr"))
		log.Info(signature)

	} else {

	}

	log.Info("XXXX")
}
