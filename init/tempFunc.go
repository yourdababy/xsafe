package init

import (
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"time"
)

func TempFuncInit() {
	_ = beego.AddFuncMap("MarkdownToHtml", MarkdownToHtml)
	_ = beego.AddFuncMap("Date", Date)
}

// Date 传递时间戳 获取 年-月-日 时:分
func Date(timeSecond int) string {
	t := time.Unix(int64(timeSecond), 0)
	return t.Format("2006-01-02 15:04")
}

// MarkdownToHtml markdown文档转html
func MarkdownToHtml(markdown string) string {
	html := blackfriday.MarkdownBasic([]byte(markdown))
	return string(html)
}