package blog

import (
	"blog/controllers/blog"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &blog.MainController{}, "get:Index")                       //首页
	beego.Router("/about", &blog.MainController{}, "get:About")                  //关于我
	beego.Router("/article", &blog.MainController{}, "get:Article")              //文章列表
	beego.Router("/article_detail", &blog.MainController{}, "get:ArticleDetail") //文章详情
	beego.Router("/board", &blog.MainController{}, "get:Board")                  //留言板
	beego.Router("/mood", &blog.MainController{}, "get:Mood")                    //碎言碎语
}
