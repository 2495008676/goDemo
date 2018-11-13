package admin

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//index.html  首页
func (c *MainController) Index() {
	c.Data["title"] = "后台管理中心"
	c.TplName = "admin/index.html"
}
