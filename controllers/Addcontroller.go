package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Package controllers 处理添加竞赛的HTTP请求和响应。
// 它包括 AddController，该控制器处理添加竞赛到数据库的 GET 和 POST 请求。

type AddController struct {
	beego.Controller // AddController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

// Get 方法处理 GET 请求，用于显示 add.html 模板。
func (c *AddController) Get() {
	c.TplName = "add.html" // 设置模板名称。
}

// Doadd 方法处理 POST 请求，用于将竞赛添加到数据库。
func (c *AddController) Doadd() {
	var competition models.Competitions                          // 定义竞赛结构体。
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &competition) // 解析请求体，将数据存储到竞赛结构体中。
	if err != nil {
		fmt.Println(err)                                                  // 如果解析失败，打印错误信息。
		c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"} // 返回 JSON 响应，指示解析失败。
		c.ServeJSON()                                                     // 将 JSON 响应发送给客户端。
		return
	}
	om := orm.NewOrm()                 // 创建 ORM 对象。
	id, err := om.Insert(&competition) // 将竞赛数据插入到数据库中。
	if id > 0 {                        // 如果插入成功。
		c.Data["json"] = LoginResponse{Success: true, Message: "增加成功"} // 返回 JSON 响应，指示插入成功。
	} else {
		c.Data["json"] = LoginResponse{Success: false, Message: "增加失败"} // 返回 JSON 响应，指示插入失败。
	}
	c.ServeJSON() // 将 JSON 响应发送给客户端。
}
