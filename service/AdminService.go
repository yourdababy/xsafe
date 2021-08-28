package service

import (
	"github.com/astaxie/beego/orm"
	"xsafe/models"
)

type AdminService struct {

}

func (this *AdminService) GetAllClassify() []models.Classify {
	classify := make([]models.Classify, 1)
	table := new(models.Classify)

	_, _ = orm.NewOrm().QueryTable(table).All(&classify)
	return classify
}
