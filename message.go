package alidayu

import "time"

const (
	MESSAGE_METHOD_DOUBLE    string = "alibaba.aliqin.fc.voice.num.doublecall" // 多方通话
	MESSAGE_METHOD_TTS              = "alibaba.aliqin.fc.tts.num.singlecall"   // 文本转语音通知
	MESSAGE_METHOD_VOICE            = "alibaba.aliqin.fc.voice.num.singlecall" // 发送语音
	MESSAGE_METHOD_SMS              = "alibaba.aliqin.fc.sms.num.send"         // 短信发送
	MESSAGE_METHOD_SMS_QUERY        = "alibaba.aliqin.fc.sms.num.query"        // 短信发送记录查询
)

const (
	MESSAGE_FORMAT_JSON string = "json" // 短信JSON参数，推荐
	MESSAGE_FORMAT_XML         = "xml"  // 短信XML参数
)

const (
	MESSAGE_SIGN_METHOD_HMAC string = "hmac" // 签名方式-HMAC
	MESSAGE_SIGN_METHOD_MD5         = "md5"  // 签名方式-MD5
)

type IMessage interface {
	Error() error
	Increase() int
}

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
	TryCount     int    `json:"-"`
}

func NewMessage(method string) *Message {
	return &Message{
		Method:     method,
		AppKey:     AppKey,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Format:     MESSAGE_FORMAT_JSON,
		Version:    "2.0",
		SignMethod: MESSAGE_SIGN_METHOD_MD5,
		Err:        nil,
		TryCount:   0,
	}
}

func (m *Message) Error() error {
	if m.Err != nil {
		return m.Err
	}
	switch m.Method {
	case MESSAGE_METHOD_DOUBLE, MESSAGE_METHOD_TTS, MESSAGE_METHOD_VOICE:
	case MESSAGE_METHOD_SMS, MESSAGE_METHOD_SMS_QUERY:
	default:
		return ErrMessageMethod
	}
	switch m.Format {
	case MESSAGE_FORMAT_JSON, MESSAGE_FORMAT_XML:
	default:
		return ErrMessageFormat
	}
	switch m.SignMethod {
	case MESSAGE_SIGN_METHOD_HMAC, MESSAGE_SIGN_METHOD_MD5:
	default:
		return ErrMessageSignMethod
	}
	return nil
}

func (m *Message) Increase() int {
	m.TryCount++
	return m.TryCount
}
