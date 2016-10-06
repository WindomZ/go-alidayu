package alidayu

import "encoding/json"

// MessageTTS 文本转语音通知
type MessageTTS struct {
	Message       `json:""`
	Param         string `json:"tts_param"`
	CalledTel     string `json:"called_num"`
	ShowCalledTel string `json:"called_show_num"`
	TTSCode       string `json:"tts_code"`
}

// NewMessageTTS 创建文本转语音通知
func NewMessageTTS() *MessageTTS {
	return &MessageTTS{
		Message: *NewMessage(MESSAGE_METHOD_TTS),
	}
}

// SetCalledTel 设置文本转语音通知的接收方电话号码
func (s *MessageTTS) SetCalledTel(tel string, show string) *MessageTTS {
	s.CalledTel = tel
	s.ShowCalledTel = show
	return s
}

// SetContent 设置文本转语音通知的内容模板code和参数param
func (s *MessageTTS) SetContent(code string, param map[string]string) *MessageTTS {
	s.TTSCode = code
	data, err := json.Marshal(param)
	if err != nil {
		s.Err = err
	} else {
		s.Param = string(data)
	}
	return s
}
