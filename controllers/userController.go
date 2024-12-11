package controllers

import (
	"encoding/json"                // 导入 json 包，用于处理 JSON 数据。
	"firstDemo/models"             // 导入 models 包，用于引用 Userinfo 结构体。
	"firstDemo/util"               // 导入自定义的 util 包，用于使用 Redis 缓存。
	"github.com/astaxie/beego"     // 导入 beego 包。
	"github.com/astaxie/beego/orm" // 导入 beego 的 ORM 包，用于数据库操作。
	"golang.org/x/crypto/bcrypt"   // 导入 bcrypt 包，用于密码哈希比较。
)

type UserController struct {
	beego.Controller // 嵌入 beego.Controller，提供 HTTP 请求处理功能
}

const (
	USERNAME = "username" // 定义常量 USERNAME 为 "username"，用于存储用户名。
)

type LoginRequest struct {
	Username string `json:"username"` // 定义 LoginRequest 结构体，用于存储登录请求中的用户名和密码。
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"` // 定义 LoginResponse 结构体，用于存储登录响应中的成功标识和消息。
	Message string `json:"message"`
}

func (c *UserController) Get() {
	c.TplName = "addUser.html" // 设置模板名称为 "addUser.html"。
}

func (c *UserController) Login() {
	var loginRequest LoginRequest                                 // 定义 LoginRequest 结构体变量，用于存储请求中的用户名和密码。
	var user models.Userinfo                                      // 定义 Userinfo 结构体变量，用于存储数据库中的用户信息。
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginRequest) // 解析请求的 JSON 数据并存储到 loginRequest 变量中。
	if err != nil {                                               // 如果解析过程中出现错误。
		c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"} // 设置响应数据为解析请求失败信息。
		c.ServeJSON()                                                     // 发送 JSON 响应给客户端。
		return                                                            // 结束函数执行。
	}

	o := orm.NewOrm()                                                                // 创建 ORM 对象。
	err = o.QueryTable("userinfo").Filter("uname", loginRequest.Username).One(&user) // 查询数据库中指定用户名的用户信息。
	if err == orm.ErrNoRows {                                                        // 如果未找到对应的用户。
		c.Data["json"] = LoginResponse{Success: false, Message: "用户不存在"} // 设置响应数据为用户不存在信息。
		c.ServeJSON()                                                    // 发送 JSON 响应给客户端。
		return                                                           // 结束函数执行。
	}

	// 比较数据库中的哈希密码与提供的密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Upwd), []byte(loginRequest.Password)); err != nil { // 比较数据库中的哈希密码与提供的密码是否匹配。
		// 密码不匹配
		c.Data["json"] = LoginResponse{Success: false, Message: "密码错误"} // 设置响应数据为密码错误信息。
		c.ServeJSON()                                                   // 发送 JSON 响应给客户端。
		return                                                          // 结束函数执行。
	}

	// 密码匹配，登录成功
	//c.SetSession(USERNAME, user.Uname) // 设置 Session 中的用户名信息。
	//Redis缓存登入信息用于是否已经是登入状态
	err = util.NewRedisCache().Set(USERNAME, user.Uname)           // 使用 Redis 缓存用户名信息。
	c.Data["json"] = LoginResponse{Success: true, Message: "登录成功"} // 设置响应数据为登录成功信息。
	c.ServeJSON()                                                  // 发送 JSON 响应给客户端。
}

func IsUserLoggedIn(c *beego.Controller) bool {
	//username := c.GetSession(USERNAME) // 从 Session 中获取用户名信息。
	username, err := util.NewRedisCache().Get(USERNAME) // 从 Redis 缓存中获取用户名信息。
	if err != nil {                                     // 如果获取过程中出现错误。
		// Handle the error, e.g., log it or return false
		beego.Error("Error getting username from Redis cache:", err) // 输出错误日志。
		return false                                                 // 返回 false。
	}
	if username == nil { // 如果用户名为空。
		return false // 返回 false。
	}
	return true // 返回 true。
}
