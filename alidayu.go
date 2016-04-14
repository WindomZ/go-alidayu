package alidayu

import "time"

var postbox *Postbox

// 初始化服务
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

// 设置服务队列
// thread_count:   同时处理线程数
// queue_capacity: 每个线程的承载数量能力
func SetAlidayuQueue(thread_count, queue_capacity int) {
	postbox = NewPostbox(thread_count, queue_capacity).Start()
}

func closeAlidayu() error {
	if postbox != nil {
		return postbox.Close()
	}
	return nil
}

// 关闭服务
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

// 设置回调
func SetCallback(callback CALLBACK) {
	if postbox != nil {
		postbox.SetCallback(callback)
	}
}

// 发送信息
func SendMessage(msg interface{}) error {
	if postbox != nil && msg != nil {
		return postbox.StuffMessage(msg)
	}
	return nil
}

// 服务是否闲置(一般用于测试)
func IsIdle() bool {
	if postbox != nil {
		return postbox.IsIdle()
	}
	return true
}
