package models

import "github.com/astaxie/beego/orm"

// Registrations 表示应用中的报名信息模型。
type Registrations struct {
	Id            int `json:"id"`             // 报名信息ID
	CompetitionId int `json:"competition_id"` // 比赛ID
	UserId        int `json:"user_id"`        // 用户ID
}

// RegistrationsToRespDesc 将报名信息转换为响应描述的方法。
func (Registration *Registrations) RegistrationsToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":            Registration.Id,
		"CompetitionId": Registration.CompetitionId,
		"UserId":        Registration.UserId,
	}

	return respInfo
}

// UpdateRegistrationResponse 表示更新报名信息的响应模型。
type UpdateRegistrationResponse struct {
	Success bool   `json:"success"` // 更新是否成功
	Message string `json:"message"` // 更新的消息
}

// GetRegistrationById 根据报名信息ID获取报名信息的方法。
func GetRegistrationById(id int) (*Registrations, error) {
	o := orm.NewOrm()
	Registration := Registrations{Id: id}
	err := o.Read(&Registration)
	if err != nil {
		return nil, err
	}
	return &Registration, nil
}
