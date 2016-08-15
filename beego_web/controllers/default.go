package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "nuaays.com"
	c.Data["Email"] = "nuaays@gmail.com"
	c.TplName = "index.tpl"
}


func (main *MainController) HelloSitepoint(){
	main.Data["Website"] = "My Website using Beego"
	main.Data["Email"] = "nuaays@gmail.com"
	main.Data["EmailName"] = "Yang Sen"
	main.TplName = "default/hello-sitepoint.tpl"
}

func (main *MainController) HelloNumber(){
	
        main.Data["Website"] = "My Website using Beego"
        main.Data["Email"] = "nuaays@gmail.com"
        main.Data["EmailName"] = "Yang Sen"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
        main.TplName = "default/hello-number.tpl"
}

