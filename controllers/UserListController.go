// Package controllers 管理应用程序的控制器逻辑。
package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
)

// UserListController 管理与用户相关的操作。
type UserListController struct {
	beego.Controller
}

// PageParam 表示分页的参数。
type PageParam struct {
	Pagesize int `json:"pagesize"` // 每页显示多少条
	Pagenum  int `json:"pagenum"`  // 第几页
}

// Page 计算总页数和当前页码。
func Page(pageSize int, pageNu int, count int64) (pageNum int) {
	pageCount := math.Ceil((float64(count) / float64(pageSize))) // 计算总页数
	fmt.Println("总页数pageCount", pageCount)
	pageNum = pageNu // 获取当前页码
	fmt.Println("获取当前页码pageNum", pageNum)
	if pageNum == 0 || pageNum == -1 { // 如果当前页码无效，则将其设置为1
		pageNum = 1
	}
	return pageNum
}

// Get 处理用户列表的 GET 请求。
func (c *UserListController) Get() {
	if !IsUserLoggedIn(&c.Controller) { // 如果用户未登录，则重定向到登录页面
		c.Redirect("/user", 302)
		return
	}
	c.TplName = "UserList.html"
}

// GetAll 检索所有用户。
func (c *UserListController) GetAll() {
	var (
		users []models.Userinfo
	)

	o := orm.NewOrm()
	o.QueryTable("userinfo").All(&users)
	if len(users) > 0 {
		var respList []interface{}
		for _, user := range users {
			respList = append(respList, user.UserToRespDesc())
			fmt.Println(user.Id, user.Uname)
		}
		c.Data["json"] = respList
		c.ServeJSON()

	}
}

// DeleteUser 处理删除用户的请求。
func (c *UserListController) DeleteUser() {
	// 从 URL 中获取用户 ID
	id, err := c.GetInt(":id")
	if err != nil { // 如果获取失败则返回错误信息
		c.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"}
		c.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建 ORM 对象
	// 调用 ORM 的 Delete 方法删除指定 ID 的用户信息
	_, err = o.Delete(&models.Userinfo{Id: id})
	if err != nil { // 如果删除失败则返回错误信息
		c.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = LoginResponse{Success: true, Message: "删除成功"}
	c.ServeJSON()
}

// ShowUserByPage 根据分页参数显示用户信息。
func (c *UserListController) ShowUserByPage() {
	// 解析请求体中的分页参数
	var pageinfo models.PageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pageinfo)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		c.ServeJSON()
		return
	}
	page := pageinfo.Pagenum
	pagesize := pageinfo.Pagesize
	// 获取指定页的用户信息
	users, err := models.GetUsers(page, pagesize)
	if err != nil {
		c.Data["json"] = LoginResponse{Success: false, Message: "获取数据失败"}
		c.ServeJSON()
		return
	}
	if len(users) > 0 { // 如果有用户信息则构建统一响应结构体并返回
		var repList []interface{}
		for _, user := range users {
			repList = append(repList, user.UserToRespDesc())
		}
		response := models.UnifiedResponse{
			Code:    200,
			Message: "获取用户成功",
			Data:    repList,
		}
		c.Data["json"] = response
	}
	c.ServeJSON()
}

func (c *UserListController) GetUserList() {
	//// 获取前端传递的分页参数
	//pageSize, _ := c.GetInt64("page_size")
	//currentPage, _ := c.GetInt64("current_page")
	var pageinfo PageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pageinfo)
	if err != nil {
		c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		c.ServeJSON()
		return
	}
	currentPage := pageinfo.Pagenum
	pageSize := pageinfo.Pagesize
	o := orm.NewOrm()
	totalCount, _ := o.QueryTable("userinfo").Count()
	fmt.Println("总数量：", totalCount)

	//totalCount, _ := models.GetUserCount()
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	//// 计算总页数
	//totalPages := totalCount / pageSize
	//if totalCount % pageSize != 0 {
	//	totalPages++
	//}

	// 查询当前页的数据
	//offset := (currentPage - 1) * pageSize
	//data, _ := models.GetUsers(offset, pageSize)
	// 确保当前页码在有效范围内
	if currentPage < 1 {
		currentPage = 1
	} else if currentPage > totalPages {
		currentPage = totalPages
	}
	var users []models.Userinfo
	o.QueryTable("userinfo").Limit(pageSize, (currentPage-1)*pageSize).All(&users)

	var repList []interface{}
	for _, user := range users {
		repList = append(repList, user.UserToRespDesc())
	}
	// 构建分页结构体
	pagination := &models.Pagination{
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Data:        repList,
	}

	// 返回分页信息给前端
	c.Data["json"] = pagination
	c.ServeJSON()
}

// ShowList 处理用户列表的显示请求。
func (c *UserListController) ShowList() {
	// 解析请求体中的分页参数
	var pageinfo PageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pageinfo)
	if err != nil {
		c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		c.ServeJSON()
		return
	}
	// 获取当前页和每页显示的数量
	currentPage := int64(pageinfo.Pagenum)
	pageSize := int64(pageinfo.Pagesize)
	// 获取用户总数
	totalCount, _ := models.GetUserCount()
	// 创建分页实例
	pagination := util.NewPagination(totalCount, currentPage, pageSize)
	// 查询当前页的数据
	offset := (pagination.CurrentPage - 1) * pagination.PageSize
	var users, _ = models.GetUserlist(offset, pagination.PageSize)
	var repList []interface{}
	// 将用户信息转换为响应描述并添加到列表中
	for _, user := range users {
		repList = append(
			repList,
			user.UserToRespDesc(),
		)
	}
	pagination.Data = repList
	// 返回分页信息给前端
	c.Data["json"] = pagination
	c.ServeJSON()
}

// GetByCustomerID 根据用户ID获取用户信息。
func (c *UserListController) GetByCustomerID() {
	// 获取URL中的用户ID
	id := c.GetString(":id")
	var search []models.Userinfo
	o := orm.NewOrm()
	// 查询用户信息
	_, err := o.QueryTable("Userinfo").Filter("id", id).All(&search)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "数据获取失败"}
		c.ServeJSON()
		return
	}
	// 检查是否找到相关数据并返回响应
	if len(search) == 0 {
		c.Data["json"] = map[string]string{"message": "未找到相关数据"}
	} else {
		c.Data["json"] = search
	}
	c.ServeJSON()
}
