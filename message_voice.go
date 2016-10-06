package alidayu

// MessageVoice 发送语音
type MessageVoice struct {
	Message       `json:""`
	CalledTel     string `json:"called_num"`
	ShowCalledTel string `json:"called_show_num"`
	VoiceCode     string `json:"voice_code"`
}

// NewMessageVoice 创建发送语音
func NewMessageVoice() *MessageVoice {
	return &MessageVoice{
		Message: *NewMessage(MESSAGE_METHOD_VOICE),
	}
}

// SetCalledTel 设置发送语音接收方的电话号码
func (s *MessageVoice) SetCalledTel(tel string, show string) *MessageVoice {
	s.CalledTel = tel
	s.ShowCalledTel = show
	return s
}
