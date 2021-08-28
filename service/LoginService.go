package service

import (
	"xsafe/models"
)

type LoginService struct {
	Admin models.Admin
}

func (this *LoginService) CheckPassword(username, password string) bool {
	admin := models.Admin{}
	user, err := admin.FindByUsername(username)
	if err != nil {
		return false
	}

	password = user.StrToPwd(password)
	if user.Password == password {
		this.Admin = admin
		return true
	}
	return false
}
