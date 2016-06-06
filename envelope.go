package alidayu

type MessageErrorFunc func(interface{}, int, error)

type Envelope struct {
	Message   interface{}
	TryCount  int
	ErrorFunc MessageErrorFunc
}

func NewEnvelope(msg interface{}, f MessageErrorFunc) *Envelope {
	return &Envelope{
		Message:   msg,
		TryCount:  0,
		ErrorFunc: f,
	}
}

func (e *Envelope) Increase() *Envelope {
	e.TryCount++
	return e
}

func (e *Envelope) Valid() bool {
	return e.Message != nil
}

func (e *Envelope) SuccessToSend() {
	if e.ErrorFunc != nil {
		e.ErrorFunc(e.Message, e.TryCount, nil)
	}
}

func (e *Envelope) FailToSend(err error) {
	if e.ErrorFunc != nil {
		e.ErrorFunc(e.Message, e.TryCount, err)
	}
}
