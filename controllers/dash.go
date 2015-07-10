package controllers

import (
	"github.com/astaxie/beego"
	. "maihaoguo/util"
)

type DashController struct {
	beego.Controller
}

func (d *DashController) Get() {
	Log.Info("SSSSSS")
	d.TplNames = "main.html"

	// sess_username, _ := d.GetSession("username").(string)
	// if sess_username == "" {
	// 	//do something
	// }

}

func (d *DashController) Post() {
	d.TplNames = "main.html"

	username := d.GetString("name")
	password := d.GetString("password")

	if username == "adimin" && password == "admin" {
		d.Redirect("example", 302)

	} else {
		d.Redirect("/", 302)
	}
}
