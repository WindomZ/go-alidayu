package alidayu

import (
	. "github.com/WindomZ/go-alidayu"
	"strconv"
	"testing"
)

func Test_Send_Message(t *testing.T) {
	StartAlidayu(1, 10)

	msg := NewMessageSms(testSign)
	msg.SetTel(testTel)
	msg.SetContent(testCode, map[string]string{"code": strconv.Itoa(0), "product": "【Test_Send_Message】"})

	if err := SendMessage(msg); err != nil {
		t.Error(err)
	}

	waitAlidayu(10)
	err := CloseAlidayu(10)
	if err != nil {
		t.Error(err)
	}
}

func Test_Send_Messages(t *testing.T) {
	StartAlidayu(2, 10)

	for i := 0; i < 3; i++ {
		msg := NewMessageSms(testSign)
		msg.SetTel(testTel)
		msg.SetContent(testCode, map[string]string{"code": strconv.Itoa(i), "product": "【Test_Send_Messages】"})
		if err := SendMessage(msg); err != nil {
			t.Error(err)
		}
	}

	waitAlidayu(10)
	err := CloseAlidayu(10)
	if err != nil {
		t.Error(err)
	}
}
