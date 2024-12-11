package controllers

import (
	"firstDemo/models"             // 导入 models 包，用于引用 Registrations 结构体。
	"fmt"                          // 导入 fmt 包，用于格式化输出。
	"github.com/astaxie/beego"     // 导入 beego 包。
	"github.com/astaxie/beego/orm" // 导入 beego 的 ORM 包，用于数据库操作。
)

type RegistrationsController struct {
	beego.Controller // RegistrationsController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *RegistrationsController) Get() {
	if !IsUserLoggedIn(&c.Controller) { // 如果用户未登录。
		c.Redirect("/user", 302) // 重定向到 "/user" 页面。
		return                   // 结束函数执行。
	}
	c.TplName = "Registration.html" // 设置模板名称为 "Registration.html"。
}

func (c *RegistrationsController) GetAll() {
	var registration []models.Registrations                         // 定义 Registrations 结构体切片，用于存储注册信息。
	newOrm := orm.NewOrm()                                          // 创建 ORM 对象。
	_, err := newOrm.QueryTable("registrations").All(&registration) // 查询数据库中的所有注册信息。
	fmt.Println(err)                                                // 打印错误信息。
	var listUser []interface{}                                      // 定义空接口切片，用于存储转换后的注册信息。
	if len(registration) > 0 {                                      // 如果注册信息切片不为空。
		for _, registration := range registration { // 遍历注册信息切片。
			listUser = append(listUser, registration.RegistrationsToRespDesc()) // 将每个注册信息转换为响应描述并添加到切片中。
		}
	}
	c.Data["json"] = listUser // 设置响应数据为转换后的注册信息。
	c.ServeJSON()             // 发送 JSON 响应给客户端。
}

func (c *RegistrationsController) DeleteRegistration() {
	id, err := c.GetInt(":id") // 获取 URL 中的 id 变量的值。
	if err != nil {            // 如果获取参数过程中出现错误。
		c.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"} // 设置响应数据为获取参数失败信息。
		c.ServeJSON()                                                     // 发送 JSON 响应给客户端。
		return                                                            // 结束函数执行。
	}
	o := orm.NewOrm() // 创建 ORM 对象。
	// 调用 ORM 的 Delete 方法，&models.Registrations{Id: id} 表示删除指定 id 的注册信息。
	_, err = o.Delete(&models.Registrations{Id: id})
	if err != nil { // 如果删除过程中出现错误。
		c.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"} // 设置响应数据为删除数据失败信息。
		c.ServeJSON()                                                     // 发送 JSON 响应给客户端。
		return                                                            // 结束函数执行。
	}
	c.Data["json"] = LoginResponse{Success: true, Message: "删除成功"} // 设置响应数据为删除成功信息。
	c.ServeJSON()                                                  // 发送 JSON 响应给客户端。
}

func (c *RegistrationsController) GetByCustomerID() {
	id := c.GetString(":id")                                              // 获取 URL 中的 id 变量的值。
	var search []models.Registrations                                     // 定义 Registrations 结构体切片，用于存储查询结果。
	o := orm.NewOrm()                                                     // 创建 ORM 对象。
	_, err := o.QueryTable("Registrations").Filter("id", id).All(&search) // 根据 id 查询注册信息。
	if err != nil {                                                       // 如果查询过程中出现错误。
		c.Data["json"] = map[string]string{"error": "数据获取失败"} // 设置响应数据为数据获取失败信息。
		c.ServeJSON()                                         // 发送 JSON 响应给客户端。
		return                                                // 结束函数执行。
	}

	if len(search) == 0 { // 如果查询结果为空。
		c.Data["json"] = map[string]string{"message": "未找到相关数据"} // 设置响应数据为未找到相关数据信息。
	} else { // 如果查询结果不为空。
		c.Data["json"] = search // 设置响应数据为查询结果。
	}

	c.ServeJSON() // 发送 JSON 响应给客户端。
}
