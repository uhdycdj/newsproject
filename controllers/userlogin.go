package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsproject/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) ShowLogin() {
	c.TplName="login.html"
}

func (c *LoginController) HandleLogin() {
	name := c.GetString("userName")
	password := c.GetString("password")
    if name==""||password==""{
    	beego.Info("数据输入有误")
    	c.Redirect("/login",302)
    	return
	}
	newOrm := orm.NewOrm()
	user := models.User{}
    user.Name=name
    user.Pwd=password
	err := newOrm.Read(&user, "Name", "Pwd")
    if err!=nil{
    	beego.Info("登陆失败")
    	c.Redirect("/login",302)
    	return
	}
	c.Redirect("/index",302)
}


