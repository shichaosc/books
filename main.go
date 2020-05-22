package main

import (
	_ "I-love-reading/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/js", "static/book/dist/js")
	beego.SetStaticPath("/css", "static/book/dist/css")
	beego.SetStaticPath("/fonts", "static/book/dist/fonts")
	beego.SetStaticPath("/img", "static/book/dist/img")
	beego.Run()
}
