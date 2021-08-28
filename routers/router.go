package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"xsafe/controllers"
	"xsafe/models"
)

func init() {
    // 登录页面
	beego.Router("/", &controllers.LoginController{}, "get:Index")
	beego.Router("/do-login", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

    // 后台页面
    beego.Router("/admin", &controllers.AdminController{}, "get:Index")
    beego.Router("/admin/article-list/:id:int", &controllers.ArticleController{}, "get:List")
	beego.Router("/admin/article/import", &controllers.ArticleController{}, "get:Import")
    beego.Router("/admin/article/do-import", &controllers.ArticleController{}, "post:DoImport")
    beego.Router("/admin/article/sync", &controllers.ArticleController{}, "get:Sync")
    beego.Router("/admin/article/do-sync", &controllers.ArticleController{}, "post:DoSync")
    beego.Router("/admin/article/sync/status", &controllers.ArticleController{}, "get:GetSyncStatus")

	// 文章增删查改
	beego.Router("/admin/article/:id:int", &controllers.ArticleController{}, "get:Detail")
    beego.Router("/admin/article/:id:int/edit", &controllers.ArticleController{}, "get:Edit")
    beego.Router("/admin/article/:id:int/update", &controllers.ArticleController{}, "post:Update")
    beego.Router("/admin/article/:id:int/del", &controllers.ArticleController{}, "post:Del")

	// 校验是否登录
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("admin").(models.Admin)
		if !ok {
			ctx.Redirect(302, "/")
		}
	}
	beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)
}
