package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"xsafe/models"
	"xsafe/service"
)

type ArticleController struct {
	beego.Controller
}

// List 文章列表
func (this *ArticleController) List() {
	classifyId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	page, _ := strconv.Atoi(this.Ctx.Input.Query("page"))
	keywords := this.Ctx.Input.Query("keywords")
	if page <= 0 {
		page = 1
	}

	// 获取所有分类
	adminService := service.AdminService{}
	classify := adminService.GetAllClassify()

	// 通过分类ID 获取所有文章列表
	articleService := service.ArticleService{}
	articles, counts := articleService.GetArticles(classifyId, page, keywords)

	// 视图模板
	this.TplName = "admin/article/list.tpl"
	// 给视图发送过去的数据
	this.Data["classifyName"] = ""
	this.Data["WebSite"] = beego.AppConfig.String("website")
	this.Data["Title"] = "文章列表"
	this.Data["TeamUrl"] = beego.AppConfig.String("teamurl")
	this.Data["Classify"] = classify
	this.Data["ActiveId"] = classifyId
	this.Data["articles"] = articles
	this.Data["pager"] = articleService.Pager(counts, page, keywords)
	this.Data["counts"] = counts
	this.Data["keywords"] = keywords

	for _, m := range classify {
		if m.Id == classifyId {
			this.Data["classifyName"] = m.Name
		}
	}
}

// Detail 文章详情
func (this *ArticleController) Detail() {
	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	// 通过文章ID 获取文章详情
	articleService := service.ArticleService{}
	article := articleService.GetArticleById(articleId)

	// 获取所有分类
	adminService := service.AdminService{}
	classify := adminService.GetAllClassify()

	// 设置视图模板
	this.TplName = "admin/article/detail.tpl"
	// 给视图传递数据
	this.Data["article"] = article
	this.Data["WebSite"] = beego.AppConfig.String("website")
	this.Data["Title"] = "文章详情"
	this.Data["TeamUrl"] = beego.AppConfig.String("teamurl")
	this.Data["Classify"] = classify
	this.Data["ActiveId"] = 0
}

// Edit 编辑文章页面
func (this *ArticleController) Edit() {
	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	// 通过文章ID 获取文章详情
	articleService := service.ArticleService{}
	article := articleService.GetArticleById(articleId)

	// 获取所有分类
	adminService := service.AdminService{}
	classify := adminService.GetAllClassify()

	// 设置视图模板
	this.TplName = "admin/article/edit.tpl"
	this.Data["TeamUrl"] = beego.AppConfig.String("teamurl")
	this.Data["WebSite"] = beego.AppConfig.String("website")
	this.Data["Article"] = article
	this.Data["Classify"] = classify
	this.Data["ActiveId"] = 0
}

// Update 执行文章修改
func (this *ArticleController) Update() {
	jsonData := make(map[string]interface{})
	jsonData["code"] = 200
	jsonData["msg"] = "更新文章成功！"
	jsonData["data"] = make([]int, 0)

	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	title := this.Ctx.Input.Query("title")
	content := this.Ctx.Input.Query("content")

	// 执行修改
	articleService := service.ArticleService{}
	err := articleService.UpdateArticle(articleId, title, content)
	if err != nil {
		jsonData["code"] = 201
		jsonData["msg"] = "更新文章失败！"
	}

	this.Data["json"] = jsonData
	this.ServeJSON()
}

// Del 执行文章删除
func (this *ArticleController) Del() {
	jsonData := make(map[string]interface{})
	jsonData["code"] = 200
	jsonData["msg"] = "更新文章成功！"
	jsonData["data"] = make([]int, 0)

	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if _, err := orm.NewOrm().Delete(&models.Article{Id: articleId}); err != nil {
		jsonData["code"] = 201
		jsonData["msg"] = "删除失败，文章未找到！"
	}

	this.Data["json"] = jsonData
	this.ServeJSON()
}

// Import 文章导入
func (this *ArticleController) Import() {
	this.TplName = "admin/article/import.tpl"
}

// DoImport 执行文章导入
func (this *ArticleController) DoImport() {

}

// Sync 文章同步
func (this *ArticleController) Sync() {
	// 获取所有分类
	adminService := service.AdminService{}
	classify := adminService.GetAllClassify()

	this.TplName = "admin/article/sync.tpl"
	// 给视图发送过去的数据
	this.Data["classifyName"] = ""
	this.Data["WebSite"] = beego.AppConfig.String("website")
	this.Data["Title"] = "同步文章"
	this.Data["TeamUrl"] = beego.AppConfig.String("teamurl")
	this.Data["Classify"] = classify
	this.Data["ActiveId"] = 0
}

// DoSync 执行文章同步
func (this *ArticleController) DoSync() {
	jsonData := make(map[string]interface{})
	this.Data["json"] = jsonData
	jsonData["code"] = 200
	jsonData["msg"] = "同步数据中..."
	jsonData["data"] = make(map[string]interface{})

	articleSyncService := service.ArticleSyncService{
		Username:          this.Input().Get("username"),
		Password:          this.Input().Get("password"),
		Host:              this.Input().Get("host"),
		Port:              this.Input().Get("port"),
		Dbname:            this.Input().Get("dbname"),
		ArticleTbName:     this.Input().Get("articleTbName"),
		TitleField:        this.Input().Get("titleField"),
		ContentField:      this.Input().Get("contentField"),
		ClassifyField:     this.Input().Get("classifyField"),
		AuthorField:       this.Input().Get("authorField"),
		ClassifyTbName:    this.Input().Get("classifyTbName"),
		ClassifyIdField:   this.Input().Get("classifyIdField"),
		ClassifyNameField: this.Input().Get("classifyNameField"),
		AuthorTbName:      this.Input().Get("authorTbName"),
		AuthorIdField:     this.Input().Get("authorIdField"),
		AuthorNameField:   this.Input().Get("authorNameField"),
	}

	// 链接数据库
	err := articleSyncService.ConnDb()
	if err != nil {
		jsonData["code"] = 201
		jsonData["msg"] = "链接数据库失败！"
		this.ServeJSON()
	}
	// 关闭数据库
	defer articleSyncService.Db.Close()

	// 获取文章表数据
	articles, counts, err := articleSyncService.ReadArticleTable()
	if err != nil {
		jsonData["code"] = 201
		jsonData["msg"] = "请检查文章表信息！"
		this.ServeJSON()
	}

	// 获取分类表数据
	classify, err := articleSyncService.ReadClassifyTable()
	if err != nil {
		jsonData["code"] = 201
		jsonData["msg"] = "请检查分类表信息！"
		this.ServeJSON()
	}

	// 获取作者表数据
	authors, err := articleSyncService.ReadAuthorTable()
	if err != nil {
		jsonData["code"] = 201
		jsonData["msg"] = "请检查作者表信息！"
		jsonData["authors"] = authors
		this.ServeJSON()
	}

	// 开启协程 执行文章同步
	go articleSyncService.DoSync(articles, classify, authors)

	data := make(map[string]interface{})
	data["articles"] = articles
	data["counts"] = counts

	jsonData["data"] = data
	this.ServeJSON()
}

// GetSyncStatus 获取文章同步状态
func (this *ArticleController) GetSyncStatus() {
	articleSyncService := service.ArticleSyncService{}
	counts := articleSyncService.GetStatus()

	jsonData := make(map[string]interface{})
	this.Data["json"] = jsonData
	jsonData["code"] = 200
	jsonData["msg"] = "获取同步数量成功！"

	data := make(map[string]int)
	data["counts"] = counts

	jsonData["data"] = data
	this.ServeJSON()
}
