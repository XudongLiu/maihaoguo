package main

import (
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
	_ "maihaoguo/routers"
	. "maihaoguo/util"
)

var AccessTokenServer = mp.NewDefaultAccessTokenServer("wx8657df9e66c8277d", "2b94ba328f84216c4eb2bfa0a6cbce82", nil) // 一個應用只能有一個實例
var mpClient = mp.NewClient(AccessTokenServer, nil)

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
	Log.Info("xxxxx")

}
