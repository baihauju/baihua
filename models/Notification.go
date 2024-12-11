// models/notification.go
package models

// Notification 表示通知的结构体
type Notification struct {
	Id               int    `json:"id"`                // 通知ID
	MessageContent   string `json:"message_content"`   // 通知内容
	NotificationDate string `json:"notification_date"` // 通知日期
	IsRead           bool   `json:"is_read"`           // 通知是否已读
}

// NotificationToRespDesc 将通知结构体转换为响应信息
func (Notification *Notification) NotificationToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":                Notification.Id,               // 将通知ID加入响应信息
		"message_content":   Notification.MessageContent,   // 将通知内容加入响应信息
		"notification_date": Notification.NotificationDate, // 将通知日期加入响应信息
		"is_read":           Notification.IsRead,           // 将通知是否已读状态加入响应信息
	}

	return respInfo // 返回响应信息
}
