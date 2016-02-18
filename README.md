# go-alidayu

## Overview

* 基于最新[阿里大鱼API](http://www.alidayu.com/doc)
* 支持短信验证码、语音双呼、文本转语音通知、语音通知及短信发送记录查询
* 提供了简单轻量化的并发队列
* 方便自行封装
* 支持回调跟踪

# Getting Started

## Installation

```
go get -u github.com/WindomZ/go-alidayu
```

## Usage

```go
package main
import (
	. "github.com/WindomZ/go-alidayu"
)
func main() {
	InitAlidayu(true, "12345678", "88888888888888888888888888888888")
	StartAlidayu(1, 1)              // 设置队列大小
	msg := NewMessageSms("身份验证") // 构建验证短信
	msg.SetTel("13088888888").
		SetContent("SMS_1234567", map[string]string{"code": "1", "product": "2"})
	if err := SendMessage(msg); err != nil {
		panic(err)
	}
	`...do somethings...`
}
```

## Documentation

可以先基于 *_test.go 进行了解
(需要在`tests/init.go`中配置正确的API调用参数)

## TODO

* 语音双呼、文本转语音通知、语音通知的测试用例
* 简洁方便的文档
* 提供封装实例方案
* 更多阿里大鱼反馈支持

## LICENSE

MIT(http://opensource.org/licenses/mit-license.php)
