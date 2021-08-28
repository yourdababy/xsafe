package controllers

import (
	"github.com/astaxie/beego"
	"xsafe/service"
)

type AdminController struct {
	beego.Controller
}

// Index 后台首页
func (this *AdminController) Index() {
	// 视图模板
	this.TplName = "admin/index.tpl"

	// 获取所有分类
	adminService := service.AdminService{}
	classify := adminService.GetAllClassify()

	// 给视图发送过去的数据
	this.Data["WebSite"] = beego.AppConfig.String("website")
	this.Data["Title"] = "首页"
	this.Data["TeamUrl"] = beego.AppConfig.String("teamurl")
	this.Data["Classify"] = classify
	this.Data["ActiveId"] = 0
}
