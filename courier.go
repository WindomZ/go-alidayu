package alidayu

import "sync"

type Courier struct {
	idle  bool
	mutex *sync.Mutex
}

func NewCourier() *Courier {
	return &Courier{
		idle:  true,
		mutex: &sync.Mutex{},
	}
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

func (s *Courier) SendEnvelope(e *Envelope) {
	if e == nil {
		return
	} else if s.work() {
		go s.working(e)
	} else {
		e.FailToSend(ERR_COURIER_BUSY)
	}
}

func (s *Courier) working(e *Envelope) {
	defer s.rest()
	if ok, resp := post(e.Increase().Message); ok {
		e.SuccessToSend()
	} else {
		e.FailToSend(NewAlidayuResponseError(resp))
	}
}
