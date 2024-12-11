package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

// Userinfo 用户信息结构体
type Userinfo struct {
	Id    int    `json:"id"`    // 用户ID
	Uname string `json:"uname"` // 用户名
	Upwd  string `json:"upwd"`  // 密码
	Usex  string `json:"usex"`  // 性别
}

// PageParam 分页参数结构体
type PageParam struct {
	Pagesize int `json:"pagesize"` //每页显示多少条
	Pagenum  int `json:"pagenum"`  //第几页
}

// UnifiedResponse 统一响应结构体
type UnifiedResponse struct {
	Code    int         `json:"code"`    // 响应状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// UserToRespDesc 将Userinfo转换为响应描述
func (user *Userinfo) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":    user.Id,
		"uname": user.Uname,
		"upwd":  user.Upwd,
		"usex":  user.Usex,
	}

	return respInfo
}

// GetUsers 根据分页参数获取用户列表
func GetUsers(page, pageSize int) ([]Userinfo, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("userinfo").Count() // 查询总数量
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var users []Userinfo
	_, err := o.QueryTable("userinfo").Limit(pageSize, (page-1)*pageSize).All(&users) // 查询指定页码的用户列表
	return users, err
}

// GetUserCount 获取用户总数
func GetUserCount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("userinfo").Count() // 查询用户总数
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetUserlist 根据偏移量和限制获取用户列表
func GetUserlist(offset, limit int64) ([]Userinfo, error) {
	o := orm.NewOrm()
	var users []Userinfo
	_, err := o.QueryTable("userinfo").Offset(offset).Limit(limit).All(&users) // 查询指定偏移量和限制的用户列表
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById 根据用户ID获取用户信息
func GetUserById(id int) (*Userinfo, error) {
	o := orm.NewOrm()
	user := Userinfo{Id: id}
	err := o.Read(&user) // 根据用户ID读取用户信息
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SearchUsers 根据搜索条件查询用户
// 根据服务ID获取服务
