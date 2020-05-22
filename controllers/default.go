package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	type ret struct {
		ID   int
		Name string
	}
	c.Data["json"] = ret{
		ID:   1,
		Name: "hello",
	}

	c.ServeJSON()
}
