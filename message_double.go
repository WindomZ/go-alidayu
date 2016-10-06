package alidayu

// MessageDouble 多方通话
type MessageDouble struct {
	Message       `json:""`
	CallerTel     string `json:"caller_num"`
	ShowCallerTel string `json:"caller_show_num"`
	CalledTel     string `json:"called_num"`
	ShowCalledTel string `json:"called_show_num"`
}

// NewMessageDouble 创建多方通话
func NewMessageDouble() *MessageDouble {
	return &MessageDouble{
		Message: *NewMessage(MESSAGE_METHOD_DOUBLE),
	}
}

// SetCallerTel 设置多方通话的发起方电话号码
func (s *MessageDouble) SetCallerTel(tel string, show string) *MessageDouble {
	s.CallerTel = tel
	s.ShowCallerTel = show
	return s
}

// SetCalledTel 设置多方通话的接收方电话号码
func (s *MessageDouble) SetCalledTel(tel string, show string) *MessageDouble {
	s.CalledTel = tel
	s.ShowCalledTel = show
	return s
}
