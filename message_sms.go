package alidayu

import (
	"encoding/json"
	"strings"
)

// MessageSms 短信发送
type MessageSms struct {
	Message      `json:""`
	Type         string `json:"sms_type"`
	FreeSignName string `json:"sms_free_sign_name"`
	Param        string `json:"sms_param"`
	Tel          string `json:"rec_num"`
	TemplateCode string `json:"sms_template_code"`
}

// NewMessageSms 创建短信发送
func NewMessageSms(signName string) *MessageSms {
	return &MessageSms{
		Message:      *NewMessage(MESSAGE_METHOD_SMS),
		Type:         "normal",
		FreeSignName: signName,
	}
}

// SetTel 设置短信接收方的电话号码
func (s *MessageSms) SetTel(tel ...string) *MessageSms {
	if len(tel) == 1 {
		s.Tel = tel[0]
	} else if len(tel) > 1 {
		s.Tel = strings.Join(tel, ",")
	}
	return s
}

// SetContent 设置短信发送信息模板code及参数param
func (s *MessageSms) SetContent(code string, param map[string]string) *MessageSms {
	s.TemplateCode = code
	data, err := json.Marshal(param)
	if err != nil {
		s.Err = err
	} else {
		s.Param = string(data)
	}
	return s
}
