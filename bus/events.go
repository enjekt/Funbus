package bus

type HexabusEvent interface {
	Send(val Response)
	Receive() Response
}

type Event struct {
	responseChannel chan Response
}

func (e *Event) getChannel() chan Response {
	if e.responseChannel == nil {
		e.responseChannel = make(chan Response, 1)
	}
	return e.responseChannel
}

func (e *Event) Send(val Response) {
	e.getChannel() <- val
}

func (e *Event) Receive() Response {
	return <-e.getChannel()
}
func NewResponse(value any, err error) Response {
	return Response{value, err}
}

type Response struct {
	Value any
	Err error
}