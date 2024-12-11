package main

import (
	"firstDemo/models"
	_ "firstDemo/routers"
	"github.com/astaxie/beego"
)

func main() {

	models.Init()
	beego.Run()
}
