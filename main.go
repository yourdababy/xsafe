package main

import (
	"github.com/astaxie/beego"
	_ "xsafe/init"
	_ "xsafe/routers"
)

func main() {
	beego.Run()
}
