package alidayu

import "time"

// MessageQuery 短信发送记录查询
type MessageQuery struct {
	Message     `json:""`
	BizId       string `json:"biz_id"`
	Tel         string `json:"rec_num"`
	Date        string `json:"query_date"`
	CurrentPage int64  `json:"current_page"`
	PageSize    int64  `json:"page_size"`
}

// NewMessageQuery 创建短信发送记录查询
func NewMessageQuery() *MessageQuery {
	return &MessageQuery{
		Message:     *NewMessage(MESSAGE_METHOD_SMS_QUERY),
		Date:        time.Now().Format("20060102"),
		CurrentPage: 1,
		PageSize:    50}
}

// SetTel 设置短信发送记录查询的手机号码
func (s *MessageQuery) SetTel(tel string) *MessageQuery {
	s.Tel = tel
	return s
}

// SetPage 设置短信发送记录查询的分页参数
func (s *MessageQuery) SetPage(current, size int64) *MessageQuery {
	if current <= 0 {
		current = 1
	}
	s.CurrentPage = current
	if size <= 0 {
		size = 50
	}
	s.PageSize = size
	return s
}

// SetLast 设置只获取短信发送记录查询的最后一条发送状态
func (s *MessageQuery) SetLast() *MessageQuery {
	return s.SetPage(1, 1)
}
