package main

import (
	_ "beego_web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"beego.log"}`)
	beego.Run()
}

