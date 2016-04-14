package alidayu

import "time"

const (
	MESSAGE_METHOD_DOUBLE    string = "alibaba.aliqin.fc.voice.num.doublecall"
	MESSAGE_METHOD_TTS              = "alibaba.aliqin.fc.tts.num.singlecall"
	MESSAGE_METHOD_VOICE            = "alibaba.aliqin.fc.voice.num.singlecall"
	MESSAGE_METHOD_SMS              = "alibaba.aliqin.fc.sms.num.send"
	MESSAGE_METHOD_SMS_QUERY        = "alibaba.aliqin.fc.sms.num.query"
)

const (
	MESSAGE_FORMAT_JSON string = "json"
	MESSAGE_FORMAT_XML         = "xml"
)
const (
	MESSAGE_SIGN_METHOD_HMAC string = "hmac"
	MESSAGE_SIGN_METHOD_MD5         = "md5"
)

type Message struct {
	Method       string `json:"method"`
	AppKey       string `json:"app_key"`
	Session      string `json:"session"`
	Timestamp    string `json:"timestamp"`
	Format       string `json:"format"`
	Version      string `json:"v"`
	PartnerId    string `json:"partner_id"`
	TargetAppKey string `json:"target_app_key"`
	Simplify     bool   `json:"simplify"`
	SignMethod   string `json:"sign_method"`
	Sign         string `json:"sign"`
	Err          error  `json:"-"`
}

func EmptyMessage() *Message {
	return &Message{}
}

func NewMessage(method string) *Message {
	msg := EmptyMessage()
	msg.Method = method
	msg.AppKey = AppKey
	msg.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	msg.Format = MESSAGE_FORMAT_JSON
	msg.Version = "2.0"
	msg.SignMethod = MESSAGE_SIGN_METHOD_MD5
	return msg
}

func (s *Message) Valid() bool {
	return (s.Err == nil)
}
