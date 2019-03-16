package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsproject/models"
)

type ShowListController struct {
	beego.Controller
}

func (s *ShowListController) ShowList() {
	newOrm := orm.NewOrm()
	var articles []models.Article
	_, err := newOrm.QueryTable("Article").All(&articles)
	if err != nil {
		beego.Info("查询文章出错")
		return
	}
	s.Data["articles"] = articles
	s.TplName = "index.html"
}
