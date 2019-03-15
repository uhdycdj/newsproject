package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsproject/models"
	"path"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.html"
}

func (c *MainController) ShowIndex() {
	c.TplName="index.html"
}

func (c *MainController) ShowAdd() {
	c.TplName="add.html"
}

func (c *MainController) AddArticle() {
	articName := c.GetString("articleName")
	artiContent := c.GetString("content")

	//图片上传
	file, header, err := c.GetFile("uploadname")
	defer file.Close()
	ext := path.Ext(header.Filename)
    beego.Info(ext)
    if ext!=".jpg"&&ext!=".png"{
    	beego.Info("上传文件格式格式错误")
		c.Redirect("/addArticle",302)
	}

	if header.Size>50000000{
		beego.Info("上传文件过大")
		c.Redirect("/addArticle",302)
	}

	//对文件重命名
	filename:=time.Now().Format("2006-01-02 15:04:05")+ext
	if err!=nil{
		beego.Info("上传文件失败",err)
		c.Redirect("/addArticle",302)
	}else {
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	if articName==""||artiContent==""{
		beego.Info("添加文章数据有误")
		c.Redirect("/addArticle",302)
	}

	newOrm := orm.NewOrm()
	article := models.Article{}
    article.ArtiName=articName
    article.Acontent=artiContent
    article.Aimg="/static/img/"+filename
	_, e := newOrm.Insert(&article)
	if e!=nil{
		beego.Info("插入数据失败")
		c.Redirect("/addArticle",302)
	}
    c.Ctx.WriteString("添加文章成功!")
}

func (c *MainController) Post() {
	name := c.GetString("userName")
	password:=c.GetString("password")
	if name==""||password==""{
		beego.Info("请输入用户名和密码")
		c.Redirect("/reg",302)
	}

	newOrm := orm.NewOrm()
	user := models.User{}
	user.Name=name
	user.Pwd=password
	_, err := newOrm.Insert(&user)
	if err!=nil{
		beego.Info("插入数据失败")
		c.Redirect("/reg",302)
	}
    //c.Ctx.WriteString("注册成功!")
    c.Redirect("/login",302)
}

