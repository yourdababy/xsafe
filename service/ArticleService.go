package service

import (
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"time"
	"xsafe/models"
)

type ArticleService struct {
}

// GetArticles 查询文章 通过分类ID
func (this *ArticleService) GetArticles(classifyId, page int, keywords string) ([]*models.Article, int) {
	o := orm.NewOrm()
	article := new(models.Article)
	var articles []*models.Article

	qs := o.QueryTable(article).
		Filter("classify_id", classifyId)

	if keywords != "" {
		cond := orm.NewCondition()
		qs = qs.SetCond(cond.And("title__icontains", keywords).Or("content__icontains", keywords))
	}

	_, _ = qs.OrderBy("-id").
		Offset((page - 1) * 10).
		Limit(10).
		All(&articles)

	counts, _ := qs.Count()

	return articles, int(counts)
}

// GetArticleById 通过文章ID 获取文章详细信息
func (this *ArticleService) GetArticleById(articleId int) models.Article {
	o := orm.NewOrm()
	var article models.Article

	_ = o.QueryTable(new(models.Article)).Filter("id", articleId).One(&article)
	return article
}

// UpdateArticle 更新文章信息
func (this *ArticleService) UpdateArticle(articleId int, title, content string) error {
	o := orm.NewOrm()
	article := models.Article{Id: articleId}
	if o.Read(&article) == nil {
		article.Title = title
		article.Content = content
		_, err := o.Update(&article, "title", "content")
		return err
	}

	return nil
}

// InsertOne 写入一条数据
func (this *ArticleService) InsertOne(tmpArticle *ArticleTable) (id int64, err error) {
	o := orm.NewOrm()

	article := new(models.Article)
	article.Title = tmpArticle.Title
	article.Content = tmpArticle.Content
	article.ClassifyId = tmpArticle.ClassifyId
	article.AuthorId = tmpArticle.AuthorId
	article.CreateAt = int(time.Now().Unix())
	article.UpdateAt = int(time.Now().Unix())

	id, err = o.Insert(article)
	return
}

// Pager 返回分页数据
/*
返回一个数组，一个数组对应一个分页
[
	[
		'name' => '首页',
		'href' => '',
		'active' => true,
		'disable' => false,
	],
	[
		'name' => '上一页',
		'href' => '',
		'active' => true,
		'disable' => false,
	],
]
*/
func (this *ArticleService) Pager(counts, currentPage int, keywords string) interface{} {
	// 每页显示10条
	pageSize := 10

	// 获取分页总数
	totalPage := int(math.Ceil(float64(counts + pageSize - 1) / float64(pageSize)))

	pager := make([]map[string]interface{}, 2)
	pager[0] = map[string]interface{}{
		"name": "<i class='fa fa-arrow-left'></i>",
		"href": "?page=1" + "&keywords=" + keywords,
		"active": false,
		"disable": currentPage == 1,
	}
	pager[1] = map[string]interface{}{
		"name": "<i class='fa fa-chevron-left'></i>",
		"href": "?page=" + strconv.Itoa(currentPage - 1) + "&keywords=" + keywords,
		"active": false,
		"disable": currentPage == 1,
	}

	if totalPage > 3 && currentPage > 2{
		page := map[string]interface{}{
			"name": "...",
			"href": "",
			"active": false,
			"disable": true,
		}
		pager = append(pager, page)
	}

	for i := currentPage-1; i < totalPage; i++ {
		if i == 0 {
			continue
		}

		if i > currentPage+1 {
			break
		}
		page := make(map[string]interface{})
		page["name"] = strconv.Itoa(i)
		page["href"] = "?page=" + strconv.Itoa(i) + "&keywords=" + keywords
		page["active"] = false
		page["disable"] = false
		if i == currentPage {
			page["active"] = true
		}

		pager = append(pager, page)
	}

	if totalPage > 3 && currentPage < totalPage-2 {
		page := map[string]interface{}{
			"name": "...",
			"href": "",
			"active": false,
			"disable": true,
		}
		pager = append(pager, page)
	}

	page := map[string]interface{}{
		"name": "<i class='fa fa-chevron-right'></i>",
		"href": "?page=" + strconv.Itoa(currentPage+1) + "&keywords=" + keywords,
		"active": false,
		"disable": currentPage == totalPage-1,
	}
	pager = append(pager, page)
	page = map[string]interface{}{
		"name": "<i class='fa fa-arrow-right'></i>",
		"href": "?page=" + strconv.Itoa(totalPage-1) + "&keywords=" + keywords,
		"active": false,
		"disable": currentPage == totalPage-1,
	}
	pager = append(pager, page)

	return pager
}
