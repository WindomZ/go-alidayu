package alidayu

import "time"

var postbox *Postbox

// InitAlidayu 初始化服务
// prod:   是否应用于正式环境
// key:    alidayu账号分配给应用的AppKey
// secret: alidayu账号分配给应用的AppSecret
func InitAlidayu(prod bool, key, secret string) {
	if prod {
		AppURL = URL_PRODUCTION
	}
	AppKey = key
	AppSecret = secret
	if postbox == nil {
		SetAlidayuQueue(DEFAULT_COURIER_SIZE, DEFAULT_CAPACITY_SIZE)
	}
}

// SetAlidayuQueue 设置服务队列
// thread_count:   同时处理线程数
// queue_capacity: 每个线程的承载数量能力
func SetAlidayuQueue(thread_count, queue_capacity int) {
	postbox = NewPostbox(thread_count, queue_capacity).Start()
}

// closeAlidayu 关闭服务
func closeAlidayu() error {
	if postbox != nil {
		return postbox.Close()
	}
	return nil
}

// CloseAlidayu 关闭服务
// waitSeconds: 最大的等待时间
func CloseAlidayu(waitSeconds ...int) error {
	if len(waitSeconds) >= 1 {
		if ws := waitSeconds[0]; ws > 0 {
			for i := 0; i < ws; i++ {
				if err := closeAlidayu(); err == nil {
					return nil
				}
				time.Sleep(time.Second)
			}
		}
	}
	return closeAlidayu()
}

// IsIdle 服务是否闲置(一般用于测试)
func IsIdle() bool {
	if postbox != nil {
		return postbox.IsIdle()
	}
	return true
}

// SendMessage 默认发送信息
// 基于SetAlidayuQueue的设置，与SendMessageQueue相同
func SendMessage(msg IMessage, f ...MessageErrorFunc) error {
	return SendMessageQueue(msg)
}

// SendMessageQueue 异步发送信息
// 异步发送信息，基于SetAlidayuQueue的设置队列处理
func SendMessageQueue(msg IMessage, f ...MessageErrorFunc) error {
	if postbox != nil {
		return postbox.StuffMessage(msg, f...)
	}
	return ErrServiceClosed
}

// SendMessageDirect 同步发送信息
// 直接发送信息，尝试一次返回结果
func SendMessageDirect(msg IMessage) error {
	if msg == nil {
		return ErrMessageNil
	} else if err := msg.Error(); err != nil {
		return err
	} else if ok, resp := post(msg); !ok {
		return NewAlidayuResponseError(resp)
	}
	return nil
}
