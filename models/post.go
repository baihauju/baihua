package models

// Post 表示应用中的文章模型。
type Post struct {
	Id       int       // 文章ID
	Title    string    // 文章标题
	Content  string    // 文章内容
	Userinfo *Userinfo `orm:"rel(fk)"` // 定义外键关联字段，与Userinfo模型建立外键关系
}
