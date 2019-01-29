package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Index") //首页
	beego.Router("/about", &controllers.MainController{},"get:About") //关于我
	beego.Router("/article", &controllers.MainController{},"get:Article") //文章列表
	beego.Router("/article_detail", &controllers.MainController{},"get:ArticleDetail") //文章详情
	beego.Router("/board", &controllers.MainController{},"get:Board") //留言板
	beego.Router("/mood", &controllers.MainController{},"get:Mood") //碎言碎语
}
