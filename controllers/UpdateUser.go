package controllers

import (
	"encoding/json"                // 导入 json 包，用于处理 JSON 数据。
	"firstDemo/models"             // 导入 models 包，用于引用 Userinfo 结构体和 GetUserById 函数。
	"github.com/astaxie/beego"     // 导入 beego 包。
	"github.com/astaxie/beego/orm" // 导入 beego 的 ORM 包，用于数据库操作。
)

type UpdateUserController struct {
	beego.Controller // UpdateUserController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *UpdateUserController) Get() {
	id, _ := c.GetInt(":id")          // 获取 URL 中的 id 变量的值。
	user, _ := models.GetUserById(id) // 调用 GetUserById 函数获取指定 id 的用户信息。
	if user != nil {                  // 如果用户信息不为空。
		c.Data["user"] = user // 设置模板数据为获取的用户信息。
	}
	c.TplName = "UpdateUser.html" // 设置模板名称为 "UpdateUser.html"。
}

func (c *UpdateUserController) UpdateUser() {
	var user models.Userinfo                              // 定义 Userinfo 结构体变量，用于存储请求中的用户信息。
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user) // 解析请求的 JSON 数据并存储到 user 变量中。
	if err != nil {                                       // 如果解析过程中出现错误。
		c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"} // 设置响应数据为解析请求失败信息。
		c.ServeJSON()                                                     // 发送 JSON 响应给客户端。
		return                                                            // 结束函数执行。
	}
	o := orm.NewOrm()                                                                               // 创建 ORM 对象。
	var r orm.RawSeter                                                                              // 定义 RawSeter 变量，用于执行原生 SQL 查询。
	r = o.Raw("UPDATE userinfo SET  uname = ?,upwd=? WHERE id = ?", user.Uname, user.Upwd, user.Id) // 执行原生 SQL 更新语句。
	_, error := r.Exec()                                                                            // 执行更新操作。
	if error != nil {                                                                               // 如果更新过程中出现错误。
		c.Data["json"] = LoginResponse{Success: true, Message: "更新失败"} // 设置响应数据为更新失败信息。
	} else { // 如果更新成功。
		c.Data["json"] = LoginResponse{Success: true, Message: "更新成功"} // 设置响应数据为更新成功信息。
	}
	c.ServeJSON() // 发送 JSON 响应给客户端。
}
