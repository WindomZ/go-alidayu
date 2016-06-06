package alidayu

type MessageVoice struct {
	Message       `json:""`
	CalledTel     string `json:"called_num"`
	ShowCalledTel string `json:"called_show_num"`
	VoiceCode     string `json:"voice_code"`
}

func NewMessageVoice() *MessageVoice {
	return &MessageVoice{
		Message: *NewMessage(MESSAGE_METHOD_VOICE),
	}
}

func (s *MessageVoice) SetCalledTel(tel string, show string) *MessageVoice {
	s.CalledTel = tel
	s.ShowCalledTel = show
	return s
}
