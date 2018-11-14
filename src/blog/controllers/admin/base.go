package admin

import (
	"github.com/astaxie/beego"
	"strings"
)

type baseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
}

func (this *baseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.auth()
}

//登录状态验证
func (this *baseController) auth() {
	if this.GetSession("userInfo") == nil && this.controllerName != "account" {
		this.Redirect("/admin/account/login", 302)
	}
}
