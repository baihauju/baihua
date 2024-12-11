package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Init 用于初始化数据库连接和注册模型
func Init() {
	// 从配置文件中获取数据库连接信息
	driverName := beego.AppConfig.String("driverName")                                         // 获取数据库驱动名称
	user := beego.AppConfig.String("mysqluser")                                                // 获取数据库用户名
	pwd := beego.AppConfig.String("mysqlpwd")                                                  // 获取数据库密码
	host := beego.AppConfig.String("host")                                                     // 获取数据库主机地址
	port := beego.AppConfig.String("port")                                                     // 获取数据库端口号
	dbname := beego.AppConfig.String("dbname")                                                 // 获取数据库名称
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8" // 构建数据库连接字符串

	// 注册数据库连接
	err := orm.RegisterDataBase("default", driverName, dbConn) // 使用数据库驱动名称和连接字符串注册数据库连接
	if err != nil {
		fmt.Printf("连接数据库出错: %s", err) // 如果出现错误则打印错误信息
		return
	}
	fmt.Println("连接数据库成功") // 打印连接数据库成功的消息

	// 注册模型
	orm.RegisterModel( // 注册数据库模型
		new(Userinfo),            // 注册 Userinfo 模型
		new(Competitions),        // 注册 Competitions 模型
		new(Registrations),       // 注册 Registrations 模型
		new(CompetitionCategory), // 注册 CompetitionCategory 模型
		new(Notification),        // 注册 Notification 模型
	)

	orm.RunSyncdb("default", false, true) // 运行 Syncdb，自动创建数据库表结构
}
