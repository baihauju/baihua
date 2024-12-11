package controllers

import (
	"encoding/json"                // 导入 encoding/json 包，用于 JSON 数据的编解码。
	"firstDemo/models"             // 导入 models 包，用于引用 Userinfo 结构体。
	"fmt"                          // 导入 fmt 包，用于格式化输出。
	"github.com/astaxie/beego"     // 导入 beego 包。
	"github.com/astaxie/beego/orm" // 导入 beego 的 ORM 包，用于数据库操作。
	"golang.org/x/crypto/bcrypt"   // 导入 bcrypt 包，用于密码哈希处理。
)

type RegisterController struct {
	beego.Controller // RegisterController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *RegisterController) Get() {
	c.TplName = "register.html" // 设置模板名称为 "register.html"。
}

func (c *RegisterController) Doregister() {
	var user models.Userinfo                              // 定义 Userinfo 结构体变量，用于存储用户信息。
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user) // 解析请求的 JSON 数据到 user 变量中。
	if err != nil {                                       // 如果解析过程中出现错误。
		fmt.Println(err)                                                  // 打印错误信息。
		c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"} // 设置响应数据为解析失败信息。
		c.ServeJSON()                                                     // 发送 JSON 响应给客户端。
		return                                                            // 结束函数执行。
	}
	// 生成密码哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Upwd), 10) // 使用 bcrypt 生成密码的哈希值。
	if err != nil {                                                 // 如果生成哈希过程中出现错误。
		fmt.Println(err) // 打印错误信息。
	}
	user.Upwd = string(hash) // 将用户密码更新为哈希值。
	om := orm.NewOrm()       // 创建 ORM 对象。
	// 将数据插入数据库
	id, err := om.Insert(&user) // 将用户信息插入到数据库中。
	if id > 0 {                 // 如果插入成功。
		c.Data["json"] = LoginResponse{Success: true, Message: "注册成功"} // 设置响应数据为注册成功信息。
	} else { // 如果插入失败。
		c.Data["json"] = LoginResponse{Success: false, Message: "注册失败"} // 设置响应数据为注册失败信息。
	}
	c.ServeJSON() // 发送 JSON 响应给客户端。
}
