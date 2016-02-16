package alidayu

import (
	. "github.com/WindomZ/go-alidayu"
	"time"
)

const (
	testKey    string = "12345678"                        // TOP分配给应用的AppKey
	testSecret        = "888888aaa8a8aa8a8a8a8a8aa8a8a8a" // TOP分配给应用的AppSecret
	testTel           = "13088888888"                     // 接收手机号码
	testSign          = "身份验证"                            // 短信签名
	testCode          = "SMS_1234567"                     // 短信模板ID
)

func init() {
	InitAlidayu(true, testKey, testSecret)
}

func waitAlidayu(maxWaitSeconds int) {
	for i := 0; i < maxWaitSeconds; i++ {
		if IsIdle() {
			break
		}
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
}
