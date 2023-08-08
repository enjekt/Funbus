package bus


type RepositoryEventListener interface {
	OnRepositoryEvent(event <-chan RepositoryEvent)
}

type RepositoryEvent interface {
	HexabusEvent
}

type RepositoryRequestEvent interface {
	RepositoryEvent
}
type RepositoryResponseEvent interface {
	RepositoryEvent
}

