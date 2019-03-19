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

func (c *MainController) ShowAdd() {
	session := c.GetSession("userName")
	newOrm := orm.NewOrm()
	var articleType []models.ArticleType
	_, err := newOrm.QueryTable("ArticleType").All(&articleType)
	if err != nil {
		beego.Info("获取类型出错")
		return
	}
	c.Data["user"] = session
	c.Layout="layout.html"
	c.Data["articleType"] = articleType
	c.TplName = "add.html"
}

func (c *MainController) HandleDelete() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取id出错")
		return
	}
	newOrm := orm.NewOrm()
	article := models.Article{Id: id}
	read := newOrm.Read(&article)
	if read != nil {
		beego.Info("查询错误")
		return
	}
	newOrm.Delete(&article)
	c.Redirect("/article/index", 302)
}

func (c *MainController) ShowUpdate() {
	session := c.GetSession("userName")
	i, e := c.GetInt("id")
	if e != nil {
		beego.Info("获取文章id错误", e)
		return
	}

	newOrm := orm.NewOrm()
	article := models.Article{Id: i}
	read := newOrm.Read(&article)
	if read != nil {
		beego.Info("查询错误", read)
		return
	}
	c.Data["user"] = session
	c.Layout="layout.html"
	c.Data["article"] = article
	c.TplName = "update.html"

}

func (c *MainController) HandleUpdate() {
	id, _ := c.GetInt("id")
	artiname := c.GetString("articleName")
	content := c.GetString("content")
	file, header, err := c.GetFile("uploadname")

	//没更新图片就不用上传图片了
	if header == nil && file == nil {
		if artiname == "" || content == "" {
			beego.Info("更新数据获取失败")
			return
		}
		newOrm := orm.NewOrm()
		article := models.Article{Id: id}
		read := newOrm.Read(&article)
		if read != nil {
			beego.Info("查询数据失败")
			return
		}
		article.ArtiName = artiname
		article.Acontent = content
		_, e := newOrm.Update(&article, "ArtiName", "Acontent")
		if e != nil {
			beego.Info("更新数据发生错误")
			return
		}
		c.Redirect("/article/index", 302)
	} else {
		defer file.Close()
		ext := path.Ext(header.Filename)
		beego.Info(ext)
		if ext != ".jpg" && ext != ".png" {
			beego.Info("上传的文件格式有误")
			return
		}
		if header.Size > 40000000 {
			beego.Info("上传文件过大")
			return
		}

		filename := time.Now().Format("2006-01-02") + ext
		if err != nil {
			beego.Info("上传文件失败:", err)
		} else {
			c.SaveToFile("uploadname", "./static/img/"+filename)
		}
		if artiname == "" || content == "" {
			beego.Info("更新数据获取失败")
			return
		}
		newOrm := orm.NewOrm()
		article := models.Article{Id: id}
		read := newOrm.Read(&article)
		if read != nil {
			beego.Info("查询数据失败")
			return
		}
		article.ArtiName = artiname
		article.Acontent = content
		article.Aimg = "./static/img/" + filename
		_, e := newOrm.Update(&article, "ArtiName", "Acontent", "Aimg")
		if e != nil {
			beego.Info("更新数据发生错误")
			return
		}
		c.Redirect("/article/index", 302)
	}
}

func (c *MainController) ShowContent() {
	session := c.GetSession("userName")
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取文章id错误", err)
		return
	}

	newOrm := orm.NewOrm()
	article := models.Article{Id: id}
	e := newOrm.Read(&article)
	if e != nil {
		beego.Info("查询错误", e)
		return
	}
	article.Acount+=1
	m2M := newOrm.QueryM2M(&article, "User")
	user := models.User{Name: session.(string)}
    newOrm.Read(&user,"Name")
	m2M.Add(user)
	newOrm.Update(&article)
	var users []models.User
	newOrm.QueryTable("User").Filter("Article__Article__Id",id).Distinct().All(&users)
	c.Data["user"] = session
	c.Data["users"]=users
	c.Layout="layout.html"
	c.Data["article"] = article
	c.TplName = "content.html"
}

func (c *MainController) AddArticle() {
	articName := c.GetString("articleName")
	artiContent := c.GetString("content")
	//图片上传
	file, header, err := c.GetFile("uploadname")
	//如果没有选择图片就不上传了
	if file == nil && header == nil {
		newOrm := orm.NewOrm()
		id, err := c.GetInt("select")
		if err!=nil{
			beego.Info("选择类别失败")
			return
		}
		articleType := models.ArticleType{Id: id}
		newOrm.Read(&articleType)
		if articName == "" || artiContent == "" {
			beego.Info("添加文章数据有误")
			c.Redirect("/article/addArticle", 302)
		}
		article := models.Article{}
		article.ArtiName = articName
		article.Acontent = artiContent
		article.ArticleType=&articleType
		//article.Aimg = "/static/img/" + filename
		_, e := newOrm.Insert(&article)
		if e != nil {
			beego.Info("插入数据失败")
			c.Redirect("/article/addArticle", 302)
		}
		//c.Ctx.WriteString("添加文章成功!")
		c.Redirect("/article/index", 302)
	} else {
		defer file.Close()
		ext := path.Ext(header.Filename)
		beego.Info(ext)
		if ext != ".jpg" && ext != ".png" {
			beego.Info("上传文件格式格式错误")
			c.Redirect("/article/addArticle", 302)
		}

		if header.Size > 50000000 {
			beego.Info("上传文件过大")
			c.Redirect("/article/addArticle", 302)
		}

		//对文件重命名
		filename := time.Now().Format("2006-01-02 15:04:05") + ext
		if err != nil {
			beego.Info("上传文件失败", err)
			c.Redirect("/article/addArticle", 302)
		} else {
			c.SaveToFile("uploadname", "./static/img/"+filename)
		}
		newOrm := orm.NewOrm()
		id, err := c.GetInt("select")
		if err!=nil{
			beego.Info("选择类别失败")
			return
		}
		articleType := models.ArticleType{Id: id}
		newOrm.Read(&articleType)
		if articName == "" || artiContent == "" {
			beego.Info("添加文章数据有误")
			c.Redirect("/article/addArticle", 302)
		}
		article := models.Article{}
		article.ArtiName = articName
		article.Acontent = artiContent
		article.ArticleType=&articleType
		article.Aimg = "/static/img/" + filename
		_, e := newOrm.Insert(&article)
		if e != nil {
			beego.Info("插入数据失败")
			c.Redirect("/article/addArticle", 302)
		}
		//c.Ctx.WriteString("添加文章成功!")
		c.Redirect("/article/index", 302)
	}

}

func (c *MainController) Post() {
	name := c.GetString("userName")
	password := c.GetString("password")
	if name == "" || password == "" {
		beego.Info("请输入用户名和密码")
		c.Redirect("/reg", 302)
	}

	newOrm := orm.NewOrm()
	user := models.User{}
	user.Name = name
	user.Pwd = password
	_, err := newOrm.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败")
		c.Redirect("/reg", 302)
	}
	//c.Ctx.WriteString("注册成功!")
	c.Redirect("/login", 302)
}
