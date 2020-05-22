package models

type Book struct {
	Id              int
	TagId           int    `orm:"size(11)"`
	AuthorId        int    `orm:"size(11)"`
	Name            string `orm:"size(50)"`
	Description     string `orm:"type(text)"`
	PreviewSmallUrl string `orm:"size(200)"`
	PreviewLargeUrl string `orm:"size(200)"`
	PreviewColor    string `orm:"size(50)"`
}
