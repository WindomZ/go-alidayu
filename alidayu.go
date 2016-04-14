package alidayu

import "time"

var postbox *Postbox

func InitAlidayu(prod bool, key, secret string) {
	if prod {
		AppURL = URL_PRODUCTION
	}
	AppKey = key
	AppSecret = secret
}

func StartAlidayu(thread_count, queue_capacity int) {
	postbox = NewPostbox(thread_count, queue_capacity).Start()
}

func CloseAlidayu(waitSeconds int) (err error) {
	if postbox != nil {
		if waitSeconds > 0 {
			for i := 0; i < waitSeconds; i++ {
				if err = postbox.Close(); err == nil {
					return nil
				}
				time.Sleep(time.Second)
			}
		}
		return postbox.Close()
	}
	return nil
}

func SetCallback(callback CALLBACK) {
	if postbox != nil {
		postbox.SetCallback(callback)
	}
}

func SendMessage(msg interface{}) error {
	if postbox != nil && msg != nil {
		return postbox.StuffMessage(msg)
	}
	return nil
}

func IsIdle() bool {
	if postbox != nil {
		return postbox.IsIdle()
	}
	return true
}
