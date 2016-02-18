package alidayu

import (
	"fmt"
	. "github.com/WindomZ/go-alidayu"
	"strconv"
	"testing"
)

// 构建回调接口
type CallBack struct {
}

func (cb *CallBack) CALLBACK_Check(msg interface{}) bool {
	println(fmt.Sprintf("CALLBACK_Check: %#v", msg))
	return true
}

func (cb *CallBack) CALLBACK_Request(msg interface{}) bool {
	println(fmt.Sprintf("CALLBACK_Request: %#v", msg))
	return true
}

func (cb *CallBack) CALLBACK_Response(msg interface{}, ok bool, result string) {
	println(fmt.Sprintf("CALLBACK_Response: %#v", msg))
	println(fmt.Sprintf("CALLBACK_Response: %v", ok))
	println(fmt.Sprintf("CALLBACK_Response: %v", result))
}

func (cb *CallBack) CALLBACK_Close() {
	println("CALLBACK_Close")
}

// 测试单条短信发送
func Test_Send_Message(t *testing.T) {
	StartAlidayu(1, 10)
	SetCallback(&CallBack{}) // 设置回调接口

	msg := NewMessageSms(testSign)
	msg.SetTel(testTel)
	msg.SetContent(testCode, map[string]string{"code": strconv.Itoa(0), "product": "【Test_Send_Message】"})

	if err := SendMessage(msg); err != nil {
		t.Error(err)
	}

	wait(10)
	err := CloseAlidayu(10)
	if err != nil {
		t.Error(err)
	}
}

// 测试多条短信发送
func Test_Send_Messages(t *testing.T) {
	StartAlidayu(2, 10)

	for i := 0; i < 3; i++ { // 可以修改发送条数
		msg := NewMessageSms(testSign)
		msg.SetTel(testTel)
		msg.SetContent(testCode, map[string]string{"code": strconv.Itoa(i), "product": "【Test_Send_Messages】"})
		if err := SendMessage(msg); err != nil {
			t.Error(err)
		}
	}

	wait(10)
	err := CloseAlidayu(10)
	if err != nil {
		t.Error(err)
	}
}
