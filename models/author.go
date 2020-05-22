package models

type Author struct {
	// 当模型定义里没有主键时，符合INT类型且名称为 Id 的 Field 将被视为自增健
	Id   int
	Name string `orm:"size(50)"`
}
