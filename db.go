package main

import (
	"I-love-reading/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func sync() {
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true

	e := orm.RegisterDriver("mysql", orm.DRMySQL)

	e = orm.RegisterDataBase(name, "mysql", "apple:apple@tcp(144.34.192.223:3306)/book?charset=utf8")

	orm.RegisterModel(new(models.Chapter))
	orm.RegisterModel(new(models.Book))

	if e != nil {
		log.Fatal(e)
	}

	// 自动建表
	//noinspection ALL
	e = orm.RunSyncdb(name, force, verbose)
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	sync()
}
