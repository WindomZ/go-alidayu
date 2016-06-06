package alidayu

import "encoding/json"

type MessageTTS struct {
	Message       `json:""`
	Param         string `json:"tts_param"`
	CalledTel     string `json:"called_num"`
	ShowCalledTel string `json:"called_show_num"`
	TTSCode       string `json:"tts_code"`
}

func NewMessageTTS() *MessageTTS {
	return &MessageTTS{
		Message: *NewMessage(MESSAGE_METHOD_TTS),
	}
}

func (s *MessageTTS) SetCalledTel(tel string, show string) *MessageTTS {
	s.CalledTel = tel
	s.ShowCalledTel = show
	return s
}

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
