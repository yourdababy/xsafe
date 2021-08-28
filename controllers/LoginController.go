package controllers

import (
	"github.com/astaxie/beego"
	"xsafe/service"
)

type LoginController struct {
	beego.Controller
}

// Index 登录页面视图页
func (this *LoginController) Index() {
	// 视图模板
	this.TplName = "login/login.tpl"

	// 给视图发送过去的数据
	this.Data["WebSite"] = beego.AppConfig.String("website")
}

// DoLogin 执行登录页
func (this *LoginController) DoLogin() {
	username := this.GetString("username")
	password := this.GetString("password")

	// 验证账号密码是否正确
	loginService := service.LoginService{}
	isError := loginService.CheckPassword(username, password)

	if isError {
		// 分配Session
		this.SetSession("admin", loginService.Admin)

		// 跳转后台首页
		this.Redirect("/admin", 302)
	}

	// 账号密码不正确，重新进入登录页面
	this.Redirect("/", 302)
}

// Logout 退出登录,清空Session
func (this *LoginController) Logout() {
	this.DestroySession()

	this.Redirect("/", 302)
}
