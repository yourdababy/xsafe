package models

import "github.com/astaxie/beego/orm"

type Author struct {
	Id int
	Name string
	TmpId int `orm:"column(tmpId)"`
}

func init() {
	orm.RegisterModel(new(Author))
}

func (this *Author) TableName() string {
	return "author"
}

