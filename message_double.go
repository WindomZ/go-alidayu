package alidayu

import ()

type MessageDouble struct {
	Message       `json:""`
	CallerTel     string `json:"caller_num"`
	ShowCallerTel string `json:"caller_show_num"`
	CalledTel     string `json:"called_num"`
	ShowCalledTel string `json:"called_show_num"`
}

func NewMessageDouble() *MessageDouble {
	return &MessageDouble{Message: *NewMessage(MESSAGE_METHOD_DOUBLE)}
}

func (s *MessageDouble) SetCallerTel(tel string, show string) *MessageDouble {
	s.CallerTel = tel
	s.ShowCallerTel = show
	return s
}

func (s *MessageDouble) SetCalledTel(tel string, show string) *MessageDouble {
	s.CalledTel = tel
	s.ShowCalledTel = show
	return s
}
