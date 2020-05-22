package tests

import (
	"I-love-reading/models"
	"I-love-reading/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func init() {
	if e := orm.RegisterDriver("mysql", orm.DRMySQL); e != nil {
		fmt.Println(e)
	}

	if e := orm.RegisterDataBase(
		"default",
		"mysql",
		"apple:apple@tcp(localhost:3306)/book?charset=utf8"); e != nil {
		fmt.Println(e)
	}
}

func init() {
	orm.RegisterModel(new(models.Book), new(models.Chapter), new(models.Tag), new(models.Author))
}

func TestORM(t *testing.T) {
	o := orm.NewOrm()
	book := models.Book{Id: 3}
	err := o.Read(&book)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(book)
	}
}

func TestGetAllTag(t *testing.T) {
	ts, i, e := utils.GetAllTag()
	fmt.Println(ts)
	fmt.Println(i)
	fmt.Println(e)
}

func TestGetChapterList(t *testing.T) {
	chapters := utils.GetChapterList(2)
	fmt.Println(chapters)
}
