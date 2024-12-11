package controllers

import (
	"encoding/json"                // 导入 json 包，用于处理 JSON 数据。
	"firstDemo/models"             // 导入 models 包，用于引用 Competition 结构体和 GetCompetitionById 函数。
	"github.com/astaxie/beego"     // 导入 beego 包。
	"github.com/astaxie/beego/orm" // 导入 beego 的 ORM 包，用于数据库操作。
)

type UpdatecompetitionsController struct {
	beego.Controller // UpdatecompetitionsController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *UpdatecompetitionsController) Get() {
	id, _ := c.GetInt(":id")                        // 获取 URL 中的 id 变量的值。
	Competition, _ := models.GetCompetitionById(id) // 调用 GetCompetitionById 函数获取指定 id 的比赛信息。
	if Competition != nil {                         // 如果比赛信息不为空。
		c.Data["Competition"] = Competition // 设置模板数据为获取的比赛信息。
	}
	c.TplName = "updatecompetitions.html" // 设置模板名称为 "updatecompetitions.html"。
}

func (c *UpdatecompetitionsController) UpdateCompetition() {
	var Competition models.Competitions                          // 定义 Competition 结构体变量，用于存储请求中的比赛信息。
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Competition) // 解析请求的 JSON 数据并存储到 Competition 变量中。
	if err != nil {                                              // 如果解析过程中出现错误。
		c.Data["json"] = UpdateResponse{Success: false, Message: "解析请求失败"} // 设置响应数据为解析请求失败信息。
		c.ServeJSON()                                                      // 发送 JSON 响应给客户端。
		return                                                             // 结束函数执行。
	}
	o := orm.NewOrm()                  // 创建 ORM 对象。
	_, error := o.Update(&Competition) // 调用 ORM 的 Update 方法更新比赛信息。
	if error != nil {                  // 如果更新过程中出现错误。
		c.Data["json"] = UpdateResponse{Success: false, Message: "更新失败"} // 设置响应数据为更新失败信息。
	} else { // 如果更新成功。
		c.Data["json"] = UpdateResponse{Success: true, Message: "更新成功"} // 设置响应数据为更新成功信息。
	}
	c.ServeJSON() // 发送 JSON 响应给客户端。
}
