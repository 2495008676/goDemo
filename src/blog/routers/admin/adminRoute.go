package admin

import (
	"blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &admin.MainController{}, "get:Index")                           //后台首页
	beego.Router("/admin/index", &admin.MainController{}, "get:Index")                     //后台首页
	beego.Router("/admin/account/login", &admin.AccountController{}, "get:Login")          //登陆
	beego.Router("/admin/account/loginAjax", &admin.AccountController{}, "post:LoginAjax") //登陆
	beego.Router("/admin/account/logout", &admin.AccountController{}, "get:Logout")        //登出
}
