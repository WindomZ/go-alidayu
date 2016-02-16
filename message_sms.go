package alidayu

import (
	"encoding/json"
	"strings"
)

type MessageSms struct {
	Message      `json:""`
	Type         string `json:"sms_type"`
	FreeSignName string `json:"sms_free_sign_name"`
	Param        string `json:"sms_param"`
	Tel          string `json:"rec_num"`
	TemplateCode string `json:"sms_template_code"`
}

func NewMessageSms(signName string) *MessageSms {
	return &MessageSms{Message: *NewMessage(MESSAGE_METHOD_SMS), Type: "normal", FreeSignName: signName}
}

func (s *MessageSms) SetTel(tel ...string) *MessageSms {
	if len(tel) == 1 {
		s.Tel = tel[0]
	} else if len(tel) > 1 {
		s.Tel = strings.Join(tel, ",")
	}
	return s
}

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
