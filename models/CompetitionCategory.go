package models

// CompetitionCategory 表示比赛类别的结构体
type CompetitionCategory struct {
	Id           int    `json:"id"`            // 比赛类别的唯一标识
	CategoryName string `json:"category_name"` // 比赛类别的名称
}

// CompetitionCategoryToRespDesc 将 CompetitionCategory 结构体转换为接口类型的响应描述
func (Category *CompetitionCategory) CompetitionCategoryToRespDesc() interface{} {
	// 构建响应信息的 map 结构
	respInfo := map[string]interface{}{
		"id":            Category.Id,           // 将 ID 加入响应信息
		"category_name": Category.CategoryName, // 将类别名称加入响应信息
	}

	return respInfo // 返回响应信息
}
