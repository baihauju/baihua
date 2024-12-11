package controllers

import (
	"github.com/astaxie/beego" // 导入 beego 包。
)

type MainController struct {
	beego.Controller // MainController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"        // 设置 Website 数据为 "beego.me"。
	c.Data["Email"] = "astaxie@gmail.com" // 设置 Email 数据为 "astaxie@gmail.com"。
	c.TplName = "index.tpl"               // 设置模板名称为 "index.tpl"。
}
