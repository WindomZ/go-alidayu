package alidayu

type MessageErrorFunc func(interface{}, int, error)

type Envelope struct {
	Message   IMessage
	ErrorFunc MessageErrorFunc
}

func NewEnvelope(msg IMessage, f MessageErrorFunc) *Envelope {
	return &Envelope{
		Message:   msg,
		ErrorFunc: f,
	}
}

func (e *Envelope) Error() error {
	return e.Message.Error()
}

func (e *Envelope) Increase() int {
	return e.Message.Increase()
}

func (e *Envelope) SuccessToSend() {
	if e.ErrorFunc != nil {
		e.ErrorFunc(e.Message, e.Increase(), nil)
	}
}

func (e *Envelope) FailToSend(err error) {
	if e.ErrorFunc != nil {
		e.ErrorFunc(e.Message, e.Increase(), err)
	}
}
