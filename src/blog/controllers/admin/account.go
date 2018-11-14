package admin

import (
	"blog/models/admin"
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
	this.TplName = "admin/login.html"
}

func (this *AccountController) LoginAjax() {
	account := strings.TrimSpace(this.GetString("userName"))
	password := strings.TrimSpace(this.GetString("password"))

	if account == "" || password == "" {
		this.Data["json"] = map[string]interface{}{"status": 0, "msg": "参数为空"}
		this.ServeJSON()
		return
	}

	var user models.User
	user.UserName = account
	user.Password = password
	user.Email = "sdds"
	user.LoginCount = 1
	user.LastIp = "12"
	user.Insert()

	this.Ctx.WriteString(user.Email)
	this.Ctx.WriteString(user.LastIp)
}

//登出
func (this *AccountController) Logout() {
	//删除指定的session
	this.DelSession("userInfo")
	//销毁全部的session
	this.DestroySession()
	//指向登陆页面
	this.Redirect("/admin/account/login", 302)
}
