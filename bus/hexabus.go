package bus

type hexabus struct {
	repositoryListenerChannels []chan RepositoryEvent

}

type Hexabus interface {

	AddRepositoryListener(eventListener RepositoryEventListener)
	SendRepositoryEvent(event RepositoryEvent)

}

var eb = hexabus{}

func Get() Hexabus {
	return &eb
}


func (bus *hexabus) AddRepositoryListener(eventListener RepositoryEventListener) {
	listenerChannel := make(chan RepositoryEvent, 10)
	bus.repositoryListenerChannels = append(bus.repositoryListenerChannels, listenerChannel)
	go eventListener.OnRepositoryEvent(listenerChannel)
}

func (bus *hexabus) SendRepositoryEvent(event RepositoryEvent) {
	for _, channel := range bus.repositoryListenerChannels {
		channel <- event
	}
}