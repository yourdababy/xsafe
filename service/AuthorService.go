package service

import (
	"github.com/astaxie/beego/orm"
	"xsafe/models"
)

type AuthorService struct {

}

// ReadOrCreateByTmpId 通过传递进来的临时ID 获取到作者信息，如果没有则创建一个并返回创建后的ID
func (this *AuthorService) ReadOrCreateByTmpId(tmpId int, tmpAuthor *AuthorTable) int {
	author := models.Author{}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Author))

	qs.Filter("tmpId", tmpId)
	err := qs.One(&author)
	id := int64(author.Id)
	if err == orm.ErrNoRows {
		// 没有读取到数据 通过传递进来的临时作者 创建一个
		author.TmpId = tmpAuthor.Id
		author.Name = tmpAuthor.Name
		id, _ = o.Insert(&author)
	}

	return int(id)
}

func (this *AuthorService) InsertOne(author *models.Author) (int, error) {
	return 0, nil
}

func (this *AuthorService) GetOneById(id int) *models.Author {
	return nil
}
