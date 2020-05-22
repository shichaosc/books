package routers

import (
	"I-love-reading/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tag", &controllers.TagController{})
	beego.Router("/book", &controllers.BookController{})
	beego.Router("/chapter", &controllers.ChapterController{})
	beego.Router("/chapter/list", &controllers.ChapterController{}, "*:List")

	beego.Router("/shelf", &controllers.ShelfController{})
}
