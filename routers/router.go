package routers

import (
	"newsproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/reg", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/addArticle", &controllers.MainController{}, "get:ShowAdd;post:AddArticle")
	beego.Router("/index", &controllers.ShowListController{}, "get:ShowList")
	beego.Router("/content", &controllers.MainController{}, "get:ShowContent")
	beego.Router("/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/delete", &controllers.MainController{}, "get:HandleDelete")
}
