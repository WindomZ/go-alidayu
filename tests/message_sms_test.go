package alidayu

import (
	. "github.com/WindomZ/go-alidayu"
	"strconv"
	"testing"
)

// 测试单条短信发送
func Test_Send_Message(t *testing.T) {
	msg := NewMessageSms(testSign).SetTel(testTel).
		SetContent(
			testCode,
			map[string]string{
				"code":    strconv.Itoa(0),
				"product": "【Test_Send_Message】",
			},
		)

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

// 测试多条短信发送
func Test_Send_Messages(t *testing.T) {
	for i := 0; i < 3; i++ { // 可以修改发送条数
		msg := NewMessageSms(testSign).SetTel(testTel).
			SetContent(
				testCode,
				map[string]string{
					"code":    strconv.Itoa(i),
					"product": "【Test_Send_Messages】",
				},
			)
		if err := SendMessage(msg, func(msg interface{}, cnt int, err error) {
			if err != nil {
				t.Errorf("FAIL(%v): %#v > %v", cnt, msg, err)
			} else {
				t.Logf("SUCC(%v): %#v", cnt, msg)
			}
		}); err != nil {
			t.Error(err)
		}
	}

	wait(10)
	if err := CloseAlidayu(10); err != nil {
		t.Error(err)
	}
}
