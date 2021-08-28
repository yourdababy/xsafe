package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"strings"
)

type Admin struct {
	Id int
	Username string
	Password string
}

func init() {
	orm.RegisterModel(new(Admin))
}

func (this *Admin) TableName() string {
	return "admin"
}

func (this *Admin) FindByUsername(username string) (Admin, error) {
	o := orm.NewOrm()
	admin := Admin{Username: username}
	err := o.Read(&admin, "Username")
	return admin, err
}

func (this *Admin) StrToPwd(pwd string) string {
	// 进行MD5加密
	d := []byte(pwd)
	m := md5.New()
	m.Write(d)

	// 将字符串MD5加密后 转大写
	return strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
}
