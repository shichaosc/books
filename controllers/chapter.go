package controllers

import (
	"I-love-reading/utils"
	"github.com/astaxie/beego"
)

type ChapterController struct {
	beego.Controller
}

// 根据分类、名称、作者等信息获取书籍
func (c *ChapterController) Get() {
	bookId, _ := c.GetInt64("bookId", -1)
	chapterId, _ := c.GetInt64("chapterId", -1)

	if chapterId > 0 && bookId > 0 {
		content, n := utils.GetChapterContent(bookId, chapterId)

		c.Data["json"] = utils.ApiJson{
			Data:  content,
			Total: n,
		}
	} else {
		c.Data["json"] = utils.ApiJson{}
	}

	c.ServeJSON()
}

func (c *ChapterController) List() {
	bookId, _ := c.GetInt("bookId")
	chapterList := utils.GetChapterList(bookId)
	t := len(chapterList)
	c.Data["json"] = utils.ApiJson{
		Data:  chapterList,
		Total: int64(t),
	}
	c.ServeJSON()
}
