package controllers

import (
	"I-love-reading/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	t, i, e := utils.GetAllTag()
	if e != nil {
		fmt.Println(e)
		return
	}

	c.Data["json"] = utils.ApiJson{
		Data:  t,
		Total: i,
	}
	c.ServeJSON()
}
