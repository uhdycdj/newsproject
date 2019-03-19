package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsproject/models"
)

type TypeController struct {
	beego.Controller
}

func (t *TypeController) ShowAddType() {
	session := t.GetSession("userName")
	newOrm := orm.NewOrm()
	var artiTypes []models.ArticleType
	_, err := newOrm.QueryTable("ArticleType").All(&artiTypes)
	if err != nil {
		beego.Info("没有获取到类型数据")
	}
	t.Data["user"] = session
	t.Layout="layout.html"
	t.Data["articleType"] = artiTypes
	t.TplName = "addType.html"
}

func (t *TypeController) HandleAddType() {
	tname := t.GetString("typeName")
	if tname == "" {
		beego.Info("获取类型信息为空")
		//没有输入类型名称，返回到添加页面
		t.Redirect("/article/addType", 302)
		return
	}
	newOrm := orm.NewOrm()
	articleType := models.ArticleType{}
	articleType.Tname = tname
	_, err := newOrm.Insert(&articleType)
	if err != nil {
		beego.Info("插入数据失败")
		return
	}
	t.Redirect("/article/addType", 302)
}
