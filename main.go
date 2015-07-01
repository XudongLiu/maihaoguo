package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "maihaoguo/routers"
)

func main() {
	beego.Run()
	log := logs.NewLogger(10000)
}
