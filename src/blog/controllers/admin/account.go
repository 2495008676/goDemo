package admin

import (
	"strings"
)

type AccountController struct {
	baseController
}

//登录
func (this *AccountController) Login() {
	if this.GetSession("userInfo") != nil {
		this.Redirect("/admin/index", 302)
	}

	account := strings.TrimSpace(this.GetString("userName"))
	password := strings.TrimSpace(this.GetString("password"))

	if account != "" && password != "" {
		this.Redirect("/admin/index", 302)
	}
	this.TplName = "admin/login.html"
}

//登出
func (this *AccountController) Logout() {

}
