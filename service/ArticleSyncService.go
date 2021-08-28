package service

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ArticleSyncService struct {
	Username          string
	Password          string
	Host              string
	Port              string
	Dbname            string
	ArticleTbName     string
	TitleField        string
	ContentField      string
	ClassifyField     string
	AuthorField       string
	ClassifyTbName    string
	ClassifyIdField   string
	ClassifyNameField string
	AuthorTbName      string
	AuthorIdField     string
	AuthorNameField   string

	Db *sql.DB
}

var bm cache.Cache

// ArticleTable 外部文章表
type ArticleTable struct {
	Id         int
	Title      string
	Content    string
	ClassifyId int
	AuthorId   int
}

// ClassifyTable 外部分类表
type ClassifyTable struct {
	Id   int
	Name string
}

// AuthorTable 外部作者表
type AuthorTable struct {
	Id   int
	Name string
}

func (this *ArticleSyncService) ConnDb() error {
	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		this.Username, this.Password, this.Host, this.Port, this.Dbname)

	db, err := sql.Open("mysql", link)
	if err != nil {
		return err
	}

	this.Db = db
	err = this.Db.Ping()
	return err
}

func (this *ArticleSyncService) ReadArticleTable() (articles []*ArticleTable, counts int64, err error) {
	query := fmt.Sprintf("SELECT %s,%s,%s,%s,%s FROM %s",
		"id", this.TitleField, this.ContentField, this.ClassifyField, this.AuthorField, this.ArticleTbName)
	rows, err := this.Db.Query(query)
	if err != nil {
		fmt.Printf("文章表读取错误！")
		return
	}

	counts = 0

	for rows.Next() {
		article := ArticleTable{}
		err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.ClassifyId, &article.AuthorId)
		if err != nil {
			return nil, 0, err
		}
		counts += 1
		articles = append(articles, &article)
	}

	err = rows.Err()
	if err != nil {
		return nil, 0, err
	}

	return articles, counts, nil
}

// ReadClassifyTable 读取分类表
func (this *ArticleSyncService) ReadClassifyTable() (classify []*ClassifyTable, err error) {
	query := fmt.Sprintf("SELECT %s,%s FROM %s",
		this.ClassifyIdField, this.ClassifyNameField, this.ClassifyTbName)
	rows, err := this.Db.Query(query)
	if err != nil {
		fmt.Printf("分类表读取错误！")
		return
	}

	for rows.Next() {
		classifyTable := ClassifyTable{}
		err := rows.Scan(&classifyTable.Id, &classifyTable.Name)
		if err != nil {
			return nil, err
		}
		classify = append(classify, &classifyTable)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return
}

// ReadAuthorTable 读取作者表
func (this *ArticleSyncService) ReadAuthorTable() (authors []*AuthorTable, err error) {
	query := fmt.Sprintf("SELECT %s,%s FROM %s",
		this.AuthorIdField, this.AuthorNameField, this.AuthorTbName)
	rows, err := this.Db.Query(query)
	if err != nil {
		fmt.Printf("作者读取错误！")
		return
	}

	for rows.Next() {
		author := AuthorTable{}
		err := rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return
}

// DoSync 执行数据同步
func (this *ArticleSyncService) DoSync(articles []*ArticleTable, classify []*ClassifyTable, authors []*AuthorTable) {
	// 设置缓存相关信息
	// 设置每60秒 检查过期值 并清除
	bm, _ = cache.NewCache("memory", `{"interval":60}`)

	articleService := ArticleService{}
	classifyService := ClassifyService{}
	authorService := AuthorService{}

	for _, article := range articles {
		// 如果分类ID不为0 则读取分类表
		if article.ClassifyId != 0 {
			tmpClassify := ClassifyTable{
				Id:   0,
				Name: "未知",
			}

			for _, classifyItem := range classify {
				if classifyItem.Id == article.ClassifyId {
					tmpClassify.Id = classifyItem.Id
					tmpClassify.Name = classifyItem.Name
					break
				}
			}
			article.ClassifyId = classifyService.ReadOrCreateByTmpId(article.ClassifyId, &tmpClassify)
		}

		// 如果作者ID不为0 则读取作者表
		if article.AuthorId != 0 {
			tmpAuthor := AuthorTable{
				Id:   0,
				Name: "未知",
			}
			for _, authorItem := range authors {
				if authorItem.Id == article.AuthorId {
					tmpAuthor.Id = authorItem.Id
					tmpAuthor.Name = authorItem.Name
					break
				}
			}
			article.AuthorId = authorService.ReadOrCreateByTmpId(article.AuthorId, &tmpAuthor)
		}

		_, _ = articleService.InsertOne(article)

		// 同步数量+1
		this.InputStatus()
	}
}

// InputStatus 写入当前同步状态
func (this *ArticleSyncService) InputStatus() {
	counts := this.GetStatus() + 1

	// 获取之前的数量 +1 再写入
	_ = bm.Put("articleSyncCounts", counts, 10 * time.Second)
}

// GetStatus 获取当前同步状态
func (this *ArticleSyncService) GetStatus() int {
	counts := bm.Get("articleSyncCounts")
	if counts != nil {
		return counts.(int)
	}

	return 0
}
