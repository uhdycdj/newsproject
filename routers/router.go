package routers

import (
	"newsproject/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/article/*",beego.BeforeRouter,beforExecFunc)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/reg", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/article/addArticle", &controllers.MainController{}, "get:ShowAdd;post:AddArticle")
	beego.Router("/article/index", &controllers.ShowListController{}, "get:ShowList")
	beego.Router("/article/content", &controllers.MainController{}, "get:ShowContent")
	beego.Router("/article/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/article/delete", &controllers.MainController{}, "get:HandleDelete")
	beego.Router("/article/addType", &controllers.TypeController{}, "get:ShowAddType;post:HandleAddType")
	beego.Router("/logout", &controllers.LoginController{}, "get:LogOut")

}
var beforExecFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil{
		ctx.Redirect(302,"/login")
	}
}
