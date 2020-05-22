package utils

import (
	"I-love-reading/models"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sort"
	"strconv"
)

func init() {
	if e := orm.RegisterDriver("mysql", orm.DRMySQL); e != nil {
		logs.Error(e)
	}

	if e := orm.RegisterDataBase(
		"default",
		"mysql",
		"apple:apple@tcp(localhost:3306)/book?charset=utf8"); e != nil {
		logs.Error(e)
	}
}

func init() {
	orm.RegisterModel(new(models.Book), new(models.Chapter), new(models.Tag), new(models.Author))
}
func GetAllTag() (t []*models.Tag, i int64, e error) {
	o := orm.NewOrm()
	tag := new(models.Tag)
	qs := o.QueryTable(tag)
	i, e = qs.All(&t)
	return t, i, e
}

type Book struct {
	Id              int
	Name            string
	Author          string
	Tag             string
	Description     string
	PreviewSmallUrl string
	PreviewLargeUrl string
	PreviewColor    string
}

func GetBookWithTagId(tagId int64) (b []Book, n int64) {

	o := orm.NewOrm()
	orm.Debug = true
	n, err := o.Raw("select b.id, b.name as name, a.name as author, "+
		"t.name as tag, b.description as description, b.preview_small_url, b.preview_large_url,"+
		"b.preview_color from book b left join author a on b.author_id=a.id left join tag t on b.tag_id=t.id "+
		"where b.tag_id=?", tagId).QueryRows(&b)
	if err == nil {
		fmt.Println("book nums: ", n)
	}
	return
}

type Content struct {
	Name    string
	Content string
	//next
}

func GetChapterContent(bookId int64, chapterId int64) (b []Content, n int64) {

	o := orm.NewOrm()
	orm.Debug = true
	n, err := o.Raw("select name, content from chapter where book_id=? and id=?", bookId, chapterId).QueryRows(&b)
	if err == nil {
		fmt.Println("content nums: ", n)
	}
	return
}

type chapter struct {
	Id   int
	Name string
	Sort string
}

func GetChapterList(bookId int) []chapter {
	o := orm.NewOrm()
	orm.Debug = true

	var chapters []chapter
	n, err := o.Raw("select id, name, sort from chapter where book_id=?", bookId).QueryRows(&chapters)
	if err == nil {
		fmt.Println("content nums: ", n)
	}
	sort.Slice(chapters, func(i, j int) bool {
		sort1, _ := strconv.Atoi(chapters[i].Sort)
		sort2, _ := strconv.Atoi(chapters[j].Sort)
		return sort1 < sort2
	})

	return chapters
}
