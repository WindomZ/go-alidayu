package alidayu

import (
	. "github.com/WindomZ/go-alidayu"
	"testing"
)

// 测试查询短信发送信息
func Test_Send_Message_Query(t *testing.T) {
	msg := NewMessageQuery().SetTel(testTel).SetLast()
	if err := SendMessage(msg, func(msg interface{}, cnt int, err error) {
		if err != nil {
			t.Errorf("FAIL(%v): %#v > %v", cnt, msg, err)
		} else {
			t.Logf("SUCC(%v): %#v", cnt, msg)
		}
	}); err != nil {
		t.Error(err)
	}

	wait(10)
	if err := CloseAlidayu(10); err != nil {
		t.Error(err)
	}
}
