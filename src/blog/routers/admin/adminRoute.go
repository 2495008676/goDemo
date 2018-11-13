package admin

import (
	"blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &admin.MainController{}, "get:Index") //后台首页
}
