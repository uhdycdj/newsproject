package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"newsproject/models"
	"math"
)

type ShowListController struct {
	beego.Controller
}

func (s *ShowListController) ShowList() {
	session := s.GetSession("userName")
	if session==nil{
		s.Redirect("/login",302)
		return
	}
	id, _ := s.GetInt("select")
	newOrm := orm.NewOrm()
	var articles []models.Article
	var artiTypes []models.ArticleType
	_, err := newOrm.QueryTable("ArticleType").All(&artiTypes)
	if err!=nil{
		beego.Info("获取类型错误")
		return
	}
	var count int64
	//count, e := newOrm.QueryTable("Article").Count()
	count, e := newOrm.QueryTable("Article").RelatedSel("ArticleType").Count()
	if id!=0{
		count,e=newOrm.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__Id",id).Count()
	}
	if e != nil {
		beego.Info("查询出错")
		return
	}
	pageSize := 3
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	pageIndex, e1 := s.GetInt("pageIndex")
	if e1 != nil {
		pageIndex = 1
	}
	start := pageSize * (pageIndex - 1)
	if id==0{
		newOrm.QueryTable("Article").Limit(pageSize, start).RelatedSel("ArticleType").All(&articles)
	}else {
		newOrm.QueryTable("Article").Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__Id",id).All(&articles)
	}
	FirstPage := false
	if pageIndex == 1 {
		FirstPage = true
	}
	LastPage := false
	if pageIndex == int(pageCount) {
		LastPage = true
	}

	s.LayoutSections = make(map[string]string)
	s.LayoutSections["contentHead"] = "head.html"
	s.Data["user"] = session
	s.Layout="layout.html"
	s.Data["pageCount"] = pageCount
	s.Data["articleType"] = artiTypes
	s.Data["FirstPage"] = FirstPage
	s.Data["LastPage"] = LastPage
	s.Data["articles"] = articles
	s.Data["pageIndex"] = pageIndex
	s.Data["count"] = count
	s.Data["typeid"] = id
	s.TplName = "index.html"
}
