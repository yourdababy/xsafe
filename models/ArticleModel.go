package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id int
	Title string
	Content string
	ClassifyId int `orm:"column(classify_id)"`
	AuthorId int `orm:"column(author_id)"`
	CreateAt int `orm:"column(create_at)"`
	UpdateAt int `orm:"column(update_at)"`
}

func init() {
	orm.RegisterModel(new(Article))
}

func (this *Article) TableName() string {
	return "article"
}
