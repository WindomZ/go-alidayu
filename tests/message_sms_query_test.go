package alidayu

import (
	. "github.com/WindomZ/go-alidayu"
	"testing"
)

// 测试查询短信发送信息
func Test_Send_Message_Query(t *testing.T) {
	StartAlidayu(1, 1)
	SetCallback(&CallBack{t})

	msg := NewMessageQuery().SetTel(testTel).SetLast()
	if err := SendMessage(msg); err != nil {
		t.Error(err)
	}

	wait(10)
	err := CloseAlidayu(10)
	if err != nil {
		t.Error(err)
	}
}
