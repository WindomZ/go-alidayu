package alidayu

import (
	"sync"
)

type Postbox struct {
	enable   bool
	barrel   chan interface{}
	Couriers []*Courier
	Callback CALLBACK
	Size     int
	mutex    *sync.Mutex
}

func NewPostbox(thread_count, queue_capacity int) *Postbox {
	if thread_count <= 0 {
		thread_count = DEFAULT_COURIER_SIZE
	}
	if queue_capacity <= 0 {
		queue_capacity = DEFAULT_CAPACITY_SIZE
	}
	cs := make([]*Courier, 0, thread_count)
	for i := 0; i < thread_count; i++ {
		cs = append(cs, NewCourier())
	}
	return &Postbox{barrel: make(chan interface{}, queue_capacity), Couriers: cs, Size: queue_capacity, mutex: &sync.Mutex{}}
}

func (s *Postbox) Start() *Postbox {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.isRunning() {
		return s
	}
	s.enable = true
	go s.service()
	return s
}

func (s *Postbox) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.enable = false
	if !s.isIdle() {
		return ERR_POSTBOX_BUSY
	}
	s.barrel <- EmptyMessage()
	return nil
}

func (s *Postbox) SetCallback(callback CALLBACK) *Postbox {
	s.Callback = callback
	return s
}

func (s *Postbox) isRunning() bool {
	return s.enable
}

func (s *Postbox) isIdle() bool {
	return len(s.barrel) == 0
}

func (s *Postbox) IsIdle() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.isIdle()
}

func (s *Postbox) isFull() bool {
	return len(s.barrel) >= s.Size
}

func (s *Postbox) IsFull() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.isFull()
}

func (s *Postbox) StuffMessage(msg interface{}) error {
	if msg == nil {
		return ERR_MESSAGE_NIL
	}
	if s.Callback != nil && !s.Callback.CALLBACK_Check(msg) {
		return ERR_NOT_PASS_CHECK
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.isRunning() {
		return ERR_SERVICE_CLOSED
	} else if s.isFull() {
		return ERR_SERVICE_OVERFLOW
	}
	s.barrel <- msg
	return nil
}

func (s *Postbox) service() {
	for s.isRunning() {
		for _, courier := range s.Couriers {
			if !s.isRunning() {
				return
			} else if courier.IsIdle() {
				courier.SendMessage(<-s.barrel, s.Callback)
			}
		}
	}
}
