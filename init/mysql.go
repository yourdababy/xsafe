package init

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL() {
	host := beego.AppConfig.DefaultString("mysqlhost", "localhost")
	port := beego.AppConfig.DefaultInt("mysqlport", 3306)
	user := beego.AppConfig.DefaultString("mysqluser", "root")
	pwd := beego.AppConfig.DefaultString("mysqlpwd", "")
	dbname := beego.AppConfig.DefaultString("msyqldbname", "xsafe")

	// 如果是开发环境才开启调试模式
	orm.Debug = true
	if beego.AppConfig.DefaultString("runmode", "dev") != "dev" {
		orm.Debug = false
	}

	// 拼接MySQL链接
	link := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pwd, host, port, dbname)

	// Golang连接池设置
	// 设置最大空闲连接
	maxIdle := 30
	err := orm.RegisterDataBase("default", "mysql", link, maxIdle)
	if err != nil {
		fmt.Printf("数据库链接错误！错误信息：%v\n", err)
	}
}
