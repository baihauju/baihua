// controllers/addregistration.go

package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// AddRegistrationController 处理添加注册的HTTP请求和响应。
type AddRegistrationController struct {
	beego.Controller // AddRegistrationController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

// RegistrationResponse 用于表示注册操作的响应信息。
type RegistrationResponse struct {
	Success bool   `json:"success"` // 表示操作是否成功。
	Message string `json:"message"` // 包含操作结果的消息。
}

// Get 方法处理 GET 请求，用于显示 addregistration.html 模板。
func (c *AddRegistrationController) Get() {
	c.TplName = "addregistration.html" // 设置模板名称。
}

// DoAddRegistration 方法处理 POST 请求，用于将注册信息添加到数据库。
func (c *AddRegistrationController) DoAddRegistration() {
	rawRequestBody := string(c.Ctx.Input.RequestBody) // 获取原始请求体数据。
	beego.Debug("Raw Request Body:", rawRequestBody)  // 打印原始请求体数据。

	var registration models.Registrations                                          // 定义注册信息结构体。
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registration); err != nil { // 解析请求体，将数据存储到注册信息结构体中。
		beego.Error("Failed to parse request:", err)                             // 如果解析失败，记录错误信息。
		c.Data["json"] = RegistrationResponse{Success: false, Message: "解析请求失败"} // 返回 JSON 响应，指示解析失败。
		c.ServeJSON()                                                            // 将 JSON 响应发送给客户端。
		return
	}

	o := orm.NewOrm()                                  // 创建 ORM 对象。
	if _, err := o.Insert(&registration); err == nil { // 将注册信息插入到数据库中。
		c.Data["json"] = RegistrationResponse{Success: true, Message: "增加成功"} // 返回 JSON 响应，指示插入成功。
	} else {
		c.Data["json"] = RegistrationResponse{Success: false, Message: "增加失败"} // 返回 JSON 响应，指示插入失败。
	}

	c.ServeJSON() // 将 JSON 响应发送给客户端。
}
