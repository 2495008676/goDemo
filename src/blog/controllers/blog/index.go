package blog

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//index.html  首页
func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "blog/index.html"
}

//about.html  关于我
func (c *MainController) About() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "blog/about.html"
}

//article.html 学无止境（文章列表）
func (c *MainController) Article() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "blog/article.html"
}

//article_detail.html 文章详情
func (c *MainController) ArticleDetail() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "blog/article_detail.html"
}

//board.html 留言板
func (c *MainController) Board() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "blog/board.html"
}

//mood.html  碎言碎语
func (c *MainController) Mood() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "blog/mood.html"
}
