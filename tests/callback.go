package alidayu

import (
	"testing"
)

// 构建回调接口
type CallBack struct {
	t *testing.T
}

func (cb *CallBack) CALLBACK_Check(msg interface{}) bool {
	//cb.t.Logf("CALLBACK_Check: %#v", msg)
	return true
}

func (cb *CallBack) CALLBACK_Request(msg interface{}) bool {
	//cb.t.Logf("CALLBACK_Request: %#v", msg)
	return true
}

func (cb *CallBack) CALLBACK_Response(msg interface{}, ok bool, result string) {
	cb.t.Logf("CALLBACK_Response: %#v", msg)
	cb.t.Logf("CALLBACK_Response: %v", ok)
	cb.t.Logf("CALLBACK_Response: %v", result)
}

func (cb *CallBack) CALLBACK_Close() {
	cb.t.Log("CALLBACK_Close")
}
