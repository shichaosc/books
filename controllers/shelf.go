package controllers

import (
	"I-love-reading/utils"
	"github.com/astaxie/beego"
)

type ShelfController struct {
	beego.Controller
}

// 根据分类、名称、作者等信息获取书籍
func (c *ShelfController) Get() {
	tagId, _ := c.GetInt64("tagId", -1)
	if tagId > 0 {
		book, n := utils.GetBookWithTagId(tagId)

		c.Data["json"] = utils.ApiJson{
			Data:  book,
			Total: n,
		}
	} else {
		c.Data["json"] = utils.ApiJson{}
	}

	c.ServeJSON()
}
