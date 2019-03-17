package main

import (
	_ "newsproject/routers"
	"github.com/astaxie/beego"
	_ "newsproject/models"
)

func main() {
	beego.AddFuncMap("showprepage", prepage)
	beego.AddFuncMap("shownextpage", nextpage)
	beego.Run()
}

/*
上一页
 */
func prepage(pageindex int) (preIndex int) {
	preIndex = pageindex - 1
	return
}

/*
下一页
 */
func nextpage(pageindex int) (nexIndex int) {
	nexIndex = pageindex + 1
	return
}
