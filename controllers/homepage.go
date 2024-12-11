package controllers

import (
	"github.com/astaxie/beego" // 导入 beego 包。
)

type HOMEPAGEController struct {
	beego.Controller // HOMEPAGEController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *HOMEPAGEController) Get() {
	c.TplName = "HOMEPAGE.html" // 设置模板名称为 "HOMEPAGE.html"。
}
