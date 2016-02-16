package alidayu

import ()

type MessageQuery struct {
	Message     `json:""`
	BizId       string `json:"biz_id"`
	Tel         string `json:"rec_num"`
	Date        string `json:"query_date"`
	CurrentPage int    `json:"current_page"`
	PageSize    int    `json:"page_size"`
}

func NewMessageQuery() *MessageQuery {
	return &MessageQuery{Message: *NewMessage(MESSAGE_METHOD_SMS_QUERY)}
}

func (s *MessageQuery) SetTel(tel string) *MessageQuery {
	s.Tel = tel
	return s
}
