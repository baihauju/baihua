package controllers

import (
	"firstDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// CompetitionCategoryController 处理竞赛类别相关的HTTP请求和响应。
type CompetitionCategoryController struct {
	beego.Controller // CompetitionCategoryController 继承自 beego.Controller，用于处理 HTTP 请求和响应。
}

// Get 方法处理 GET 请求，用于显示 CompetitionCategory.html 模板。
func (c *CompetitionCategoryController) Get() {
	c.TplName = "CompetitionCategory.html" // 设置模板名称。
}

// GetAll 方法处理 GET 请求，用于获取所有竞赛类别信息。
func (c *CompetitionCategoryController) GetAll() {
	var Category []models.CompetitionCategory                          // 定义竞赛类别切片。
	newOrm := orm.NewOrm()                                             // 创建 ORM 对象。
	_, err := newOrm.QueryTable("competition_category").All(&Category) // 查询所有竞赛类别信息。
	fmt.Println(err)                                                   // 打印错误信息。
	var listUser []interface{}                                         // 定义接口类型切片。
	if len(Category) > 0 {                                             // 如果查询到了竞赛类别信息。
		for _, Competition := range Category { // 遍历竞赛类别信息。
			listUser = append(listUser, Competition.CompetitionCategoryToRespDesc()) // 将竞赛类别信息转换为响应描述并添加到接口类型切片中。
		}
	}
	c.Data["json"] = listUser // 设置响应数据为接口类型切片。
	c.ServeJSON()             // 将 JSON 响应发送给客户端。
}

// DeleteCompetition 方法处理 DELETE 请求，用于删除指定的竞赛类别信息。
func (c *CompetitionCategoryController) DeleteCompetition() {
	id, err := c.GetInt(":id") // 获取 URL 中的 id 变量的值。
	if err != nil {            // 如果获取失败。
		c.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"} // 返回 JSON 响应，指示获取参数失败。
		c.ServeJSON()                                                     // 将 JSON 响应发送给客户端。
		return                                                            // 结束方法执行。
	}
	o := orm.NewOrm()                                      // 创建 ORM 对象。
	_, err = o.Delete(&models.CompetitionCategory{Id: id}) // 删除指定 id 的竞赛类别信息。
	if err != nil {                                        // 如果删除失败。
		c.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"} // 返回 JSON 响应，指示删除数据失败。
		c.ServeJSON()                                                     // 将 JSON 响应发送给客户端。
		return                                                            // 结束方法执行。
	}
	c.Data["json"] = LoginResponse{Success: true, Message: "删除成功"} // 返回 JSON 响应，指示删除成功。
	c.ServeJSON()                                                  // 将 JSON 响应发送给客户端。
}
