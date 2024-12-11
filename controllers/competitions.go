package controllers

import (
	"firstDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// CompetitionController 处理竞赛相关的HTTP请求和响应。
type CompetitionController struct {
	beego.Controller // CompetitionController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

// Get 方法处理 GET 请求，用于显示 Competitions.html 模板。
func (c *CompetitionController) Get() {
	if !IsUserLoggedIn(&c.Controller) { // 如果用户未登录。
		c.Redirect("/user", 302) // 重定向到用户登录页面。
		return                   // 结束方法执行。
	}
	c.TplName = "Competitions.html" // 设置模板名称。
}

// UpdateResponse 用于表示更新响应的结构体。
type UpdateResponse struct {
	Success bool   `json:"success"` // 表示更新是否成功。
	Message string `json:"message"` // 更新的消息。
}

// GetAll 方法处理 GET 请求，用于获取所有竞赛信息。
func (c *CompetitionController) GetAll() {
	var competition []models.Competitions                         // 定义竞赛切片。
	newOrm := orm.NewOrm()                                        // 创建 ORM 对象。
	_, err := newOrm.QueryTable("competitions").All(&competition) // 查询所有竞赛信息。
	fmt.Println(err)                                              // 打印错误信息。
	var listUser []interface{}                                    // 定义接口类型切片。
	if len(competition) > 0 {                                     // 如果查询到了竞赛信息。
		for _, Competition := range competition { // 遍历竞赛信息。
			listUser = append(listUser, Competition.CompetitionToRespDesc()) // 将竞赛信息转换为响应描述并添加到接口类型切片中。
		}
	}
	c.Data["json"] = listUser // 设置响应数据为接口类型切片。
	c.ServeJSON()             // 将 JSON 响应发送给客户端。
}

// DeleteCompetition 方法处理 DELETE 请求，用于删除指定的竞赛信息。
func (c *CompetitionController) DeleteCompetition() {
	id, err := c.GetInt(":id") // 获取 URL 中的 id 变量的值。
	if err != nil {            // 如果获取失败。
		c.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"} // 返回 JSON 响应，指示获取参数失败。
		c.ServeJSON()                                                     // 将 JSON 响应发送给客户端。
		return                                                            // 结束方法执行。
	}
	o := orm.NewOrm()                               // 创建 ORM 对象。
	_, err = o.Delete(&models.Competitions{Id: id}) // 删除指定 id 的竞赛信息。
	if err != nil {                                 // 如果删除失败。
		c.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"} // 返回 JSON 响应，指示删除数据失败。
		c.ServeJSON()                                                     // 将 JSON 响应发送给客户端。
		return                                                            // 结束方法执行。
	}
	c.Data["json"] = LoginResponse{Success: true, Message: "删除成功"} // 返回 JSON 响应，指示删除成功。
	c.ServeJSON()                                                  // 将 JSON 响应发送给客户端。
}

// GetByCustomerID 方法处理 GET 请求，用于根据客户ID获取竞赛信息。
func (c *CompetitionController) GetByCustomerID() {
	id := c.GetString(":id")                                             // 获取 URL 中的 id 变量的值。
	var search []models.Competitions                                     // 定义竞赛切片。
	o := orm.NewOrm()                                                    // 创建 ORM 对象。
	_, err := o.QueryTable("Competitions").Filter("id", id).All(&search) // 根据客户ID查询竞赛信息。
	if err != nil {                                                      // 如果查询失败。
		c.Data["json"] = map[string]string{"error": "数据获取失败"} // 返回 JSON 响应，指示数据获取失败。
		c.ServeJSON()                                         // 将 JSON 响应发送给客户端。
		return                                                // 结束方法执行。
	}

	if len(search) == 0 { // 如果未找到相关数据。
		c.Data["json"] = map[string]string{"message": "未找到相关数据"} // 返回 JSON 响应，指示未找到相关数据。
	} else {
		c.Data["json"] = search // 设置响应数据为查询到的竞赛信息。
	}

	c.ServeJSON() // 将 JSON 响应发送给客户端。
}
