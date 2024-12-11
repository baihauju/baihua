package models

import "github.com/astaxie/beego/orm"

// Competitions 表示比赛的结构体
type Competitions struct {
	Id                    int    `json:"id"`                       // 比赛的唯一标识
	Title                 string `json:"title"`                    // 比赛的标题
	Description           string `json:"description"`              // 比赛的描述
	Location              string `json:"location"`                 // 比赛的地点
	StartTime             string `json:"start_time"`               // 比赛的开始时间
	EndTime               string `json:"end_time"`                 // 比赛的结束时间
	IsOpenForRegistration string `json:"is_open_for_registration"` // 比赛是否开放注册
}

// CompetitionToRespDesc 将 Competitions 结构体转换为接口类型的响应描述
func (Competition *Competitions) CompetitionToRespDesc() interface{} {
	// 构建响应信息的 map 结构
	respInfo := map[string]interface{}{
		"Id":                    Competition.Id,                    // 将 ID 加入响应信息
		"Title":                 Competition.Title,                 // 将标题加入响应信息
		"Description":           Competition.Description,           // 将描述加入响应信息
		"Location":              Competition.Location,              // 将地点加入响应信息
		"StartTime":             Competition.StartTime,             // 将开始时间加入响应信息
		"EndTime":               Competition.EndTime,               // 将结束时间加入响应信息
		"IsOpenForRegistration": Competition.IsOpenForRegistration, // 将注册状态加入响应信息
	}

	return respInfo // 返回响应信息
}

// UpdateResponse 表示更新响应的结构体
type UpdateResponse struct {
	Success bool   `json:"success"` // 更新操作是否成功
	Message string `json:"message"` // 更新操作的消息
}

// GetCompetitionById 根据比赛ID获取比赛信息
func GetCompetitionById(id int) (*Competitions, error) {
	o := orm.NewOrm()                   // 创建 ORM 对象
	Competition := Competitions{Id: id} // 创建一个 Competitions 结构体，并指定 ID
	err := o.Read(&Competition)         // 通过 ORM 读取指定 ID 的比赛信息
	if err != nil {
		return nil, err // 如果出现错误则返回错误信息
	}
	return &Competition, nil // 返回获取到的比赛信息
}
