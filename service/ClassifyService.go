package service

import (
	"github.com/astaxie/beego/orm"
	"xsafe/models"
)

type ClassifyService struct {

}

// ReadOrCreateByTmpId 通过临时ID 获取到主ID，如果没有则创建并返回新ID
func (this *ClassifyService) ReadOrCreateByTmpId(tmpId int, tmpClassify *ClassifyTable) int {
	classify := models.Classify{}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Classify))

	qs.Filter("tmpId", tmpId)
	err := qs.One(&classify)
	id := int64(classify.Id)
	if err == orm.ErrNoRows {
		// 没有读取到数据 通过传递进来的临时分类 创建一个
		classify.TmpId = tmpClassify.Id
		classify.Name = tmpClassify.Name
		id, _ = o.Insert(&classify)
	}

	return int(id)
}

func (this *ClassifyService) InsertOne(classify *models.Classify) {

}

func (this *ClassifyService) GetOneById(id int) {

}
