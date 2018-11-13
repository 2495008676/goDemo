package admin

import (
	"blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/index", &admin.MainController{}, "get:Index")            //后台首页
	beego.Router("/admin/account/login", &admin.AccountController{}, "get:Login") //后台首页
}
