// controllers/notification.go
package controllers

import (
	"firstDemo/models"             // 导入 models 包，用于引用 Notification 结构体和相关方法。
	"fmt"                          // 导入 fmt 包，用于格式化输出。
	"github.com/astaxie/beego"     // 导入 beego 包。
	"github.com/astaxie/beego/orm" // 导入 beego 的 ORM 包，用于数据库操作。
)

type NotificationController struct {
	beego.Controller // NotificationController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

func (c *NotificationController) Get() {
	c.TplName = "notifications.html" // 设置模板名称为 "notifications.html"。
}

func (c *NotificationController) GetAll() {
	var Notification []models.Notification                         // 定义 Notification 结构体切片，用于存储查询结果。
	newOrm := orm.NewOrm()                                         // 创建 ORM 对象。
	_, err := newOrm.QueryTable("notification").All(&Notification) // 查询所有通知信息。
	fmt.Println(err)                                               // 打印错误信息。

	var listUser []interface{} // 定义通用类型的切片，用于存储转换后的通知信息。
	if len(Notification) > 0 { // 如果查询结果不为空。
		for _, Notification := range Notification { // 遍历查询结果中的通知信息。
			listUser = append(listUser, Notification.NotificationToRespDesc()) // 将通知信息转换为响应描述并添加到切片中。
		}
	}
	c.Data["json"] = listUser // 设置响应数据为转换后的通知信息。
	c.ServeJSON()             // 将 JSON 响应发送给客户端。
}
