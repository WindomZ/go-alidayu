package alidayu

import (
	"sync"
)

type Courier struct {
	idle  bool
	mutex *sync.Mutex
}

func NewCourier() *Courier {
	return &Courier{idle: true, mutex: &sync.Mutex{}}
}

func (s *Courier) IsIdle() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.idle
}

func (s *Courier) work() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.idle {
		return false
	}
	s.idle = false
	return true
}

func (s *Courier) rest() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.idle = true
}

func (s *Courier) SendMessage(msg interface{}, callback CALLBACK) {
	if s.work() {
		go s.working(msg, callback)
	} else if callback != nil {
		callback.CALLBACK_Response(msg, false, ERR_COURIER_BUSY.Error())
	}
}

func (s *Courier) working(msg interface{}, callback CALLBACK) {
	defer s.rest()
	if callback == nil {
		post(msg)
	} else if callback.CALLBACK_Request(msg) {
		ok, resp := post(msg)
		callback.CALLBACK_Response(msg, ok, resp)
	}
}
