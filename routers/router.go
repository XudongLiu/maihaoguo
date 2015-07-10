package routers

import (
	"github.com/astaxie/beego"
	"maihaoguo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/main", &controllers.DashController{})

}
