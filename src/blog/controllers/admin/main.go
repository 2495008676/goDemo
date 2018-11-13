package admin

import (
	"github.com/astaxie/beego"
	"os"
)

type MainController struct {
	beego.Controller
}

//index.html  首页
func (c *MainController) Index() {
	path := os.Getenv("GOPATH")
	c.Data["title"] = "后台管理中心"
	c.Data["dd"] = path
	c.TplName = "admin/index.html"
}
