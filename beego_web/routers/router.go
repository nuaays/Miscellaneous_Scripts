package routers

import (
	"beego_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.MainController{}, "get:HelloSitepoint")
    beego.Router("/hello/:id[0-9]+", &controllers.MainController{}, "get,post:HelloNumber")
}
