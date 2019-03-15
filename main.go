package main

import (
	_ "newsproject/routers"
	"github.com/astaxie/beego"
	_"newsproject/models"

)

func main() {
	beego.Run()
}

