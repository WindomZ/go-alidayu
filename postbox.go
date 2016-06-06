package alidayu

import "sync"

type Postbox struct {
	enable   bool
	barrel   chan *Envelope
	Couriers []*Courier
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
	return &Postbox{
		barrel:   make(chan *Envelope, queue_capacity),
		Couriers: cs,
		Size:     queue_capacity,
		mutex:    &sync.Mutex{},
	}
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
		return ErrPostboxBusy
	}
	s.barrel <- nil
	return nil
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

func (s *Postbox) StuffMessage(msg IMessage, f ...MessageErrorFunc) error {
	if msg == nil {
		return ErrMessageNil
	} else if err := msg.Error(); err != nil {
		return err
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.isRunning() {
		return ErrServiceClosed
	} else if s.isFull() {
		return ErrServiceOverflow
	} else if f == nil || len(f) == 0 {
		s.barrel <- NewEnvelope(msg, nil)
	} else {
		s.barrel <- NewEnvelope(msg, f[0])
	}
	return nil
}

func (s *Postbox) service() {
	for s.isRunning() {
		for _, courier := range s.Couriers {
			if !s.isRunning() {
				return
			} else if courier.IsIdle() {
				courier.SendEnvelope(<-s.barrel)
			}
		}
	}
}
