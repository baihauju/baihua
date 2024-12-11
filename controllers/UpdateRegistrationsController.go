package controllers

import (
	"encoding/json"                // 导入 json 包，用于处理 JSON 数据。
	"firstDemo/models"             // 导入 models 包，用于引用 Registration 结构体和 GetRegistrationById 函数。
	"github.com/astaxie/beego"     // 导入 beego 包。
	"github.com/astaxie/beego/orm" // 导入 beego 的 ORM 包，用于数据库操作。
)

type UpdateRegistrationsController struct {
	beego.Controller // UpdateRegistrationsController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *UpdateRegistrationsController) Get() {
	id, _ := c.GetInt(":id")                          // 获取 URL 中的 id 变量的值。
	registration, _ := models.GetRegistrationById(id) // 调用 GetRegistrationById 函数获取指定 id 的报名信息。
	if registration != nil {                          // 如果报名信息不为空。
		c.Data["Registration"] = registration // 设置模板数据为获取的报名信息。
	}
	c.TplName = "updateRegistrations.html" // 设置模板名称为 "updateRegistrations.html"。
}

func (c *UpdateRegistrationsController) UpdateRegistration() {
	var registration models.Registrations                         // 定义 Registration 结构体变量，用于存储请求中的报名信息。
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &registration) // 解析请求的 JSON 数据并存储到 registration 变量中。
	if err != nil {                                               // 如果解析过程中出现错误。
		c.Data["json"] = models.UpdateRegistrationResponse{Success: false, Message: "解析请求失败"} // 设置响应数据为解析请求失败信息。
		c.ServeJSON()                                                                         // 发送 JSON 响应给客户端。
		return                                                                                // 结束函数执行。
	}

	o := orm.NewOrm()                // 创建 ORM 对象。
	_, err = o.Update(&registration) // 调用 ORM 的 Update 方法更新报名信息。
	if err != nil {                  // 如果更新过程中出现错误。
		c.Data["json"] = models.UpdateRegistrationResponse{Success: false, Message: "更新失败"} // 设置响应数据为更新失败信息。
	} else { // 如果更新成功。
		c.Data["json"] = models.UpdateRegistrationResponse{Success: true, Message: "更新成功"} // 设置响应数据为更新成功信息。
	}
	c.ServeJSON() // 发送 JSON 响应给客户端。
}
