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
	username := c.Ctx.GetCookie("userName")
    if username!=""{
    	c.Data["userName"]=username
    	c.Data["checked"]="checked"
	}else {
		c.Data["userName"]=""
	}
	c.TplName="login.html"
}

func (c *LoginController) LogOut() {
	c.DelSession("userName")
	c.Redirect("/login",302)
}


func (c *LoginController) HandleLogin() {
	name := c.GetString("userName")
	password := c.GetString("password")
	remember := c.GetString("remember")
	beego.Info("remember=",remember)
	if remember=="on"{
		c.Ctx.SetCookie("userName",name,200)
	}else {
		c.Ctx.SetCookie("userName",name,-1)
	}
    c.SetSession("userName",name)
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
	c.Redirect("/article/index",302)
}


