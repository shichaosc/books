package models

type Chapter struct {
	// 当模型定义里没有主键时，符合INT类型且名称为 Id 的 Field 将被视为自增健
	Id      int
	BookId  int    `orm:"size(11)"`
	Name    string `orm:"size(200)"`
	Content string `orm:"type(longtext)"`
	Sort    string `orm:"size(50)"`
}
