package main

import (
	_ "blog/routers/admin"
	_ "blog/routers/blog"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/blogStatic"] = "static/blog"
	beego.Run()
}
