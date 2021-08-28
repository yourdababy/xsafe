package models

import "github.com/astaxie/beego/orm"

type Classify struct {
	Id int
	Name string
	TmpId int `orm:"column(tmpId)"`
}

func init() {
	orm.RegisterModel(new(Classify))
}

func (this *Classify) TableName() string {
	return "classify"
}

