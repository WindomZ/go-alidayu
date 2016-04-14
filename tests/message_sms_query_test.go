package alidayu

import (
	. "github.com/WindomZ/go-alidayu"
	"testing"
)

// 测试查询短信发送信息
func Test_Send_Message_Query(t *testing.T) {
	SetCallback(&CallBack{t}) //　设置测试回调，方便调试监控

	msg := NewMessageQuery().SetTel(testTel).SetLast()
	if err := SendMessage(msg); err != nil {
		t.Error(err)
	}

	wait(10)
	if err := CloseAlidayu(10); err != nil {
		t.Error(err)
	}
}
