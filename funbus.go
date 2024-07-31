package funbus

import (
	"fmt"
	"reflect"
	"sync"
)

// Mutex locking is for short periods as most of the message handling is on the channels themselves
// and the bus is simply handing off the message to the channels it finds.
type Hexabus interface {
	Subscribe(fn interface{}) error
	Unsubscribe(fn interface{}) error
	Send(event interface{})
}
type hexabus struct {
	mu       sync.Mutex
	registry map[string]map[reflect.Value]chan interface{}
}

var bus = hexabus{registry: make(map[string]map[reflect.Value]chan interface{})}

func Subscribe(fn interface{}) error {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	handlerType := reflect.TypeOf(fn)
	//fmt.Println("Name:", reflect.ValueOf(fn))
	if !(handlerType.Kind() == reflect.Func) {
		return fmt.Errorf("%s is not of type reflect.Func", reflect.TypeOf(fn).Kind())
	}
	eventName := handlerType.In(0).Name()

	var channelMap = bus.registry[eventName]
	if channelMap == nil {
		channelMap = make(map[reflect.Value]chan interface{})
		bus.registry[eventName] = channelMap
	}

	//TODO default of 10 messages have a different Subscribe with optional int?
	handlerChannel := make(chan interface{}, 10)
	handler := eventHandler{reflect.ValueOf(fn)}
	go handler.OnEvent(handlerChannel)

	channelMap[reflect.ValueOf(fn)] = handlerChannel
	return nil
}

func Unsubscribe(fn interface{}) error {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	handlerType := reflect.TypeOf(fn)
	//fmt.Println("Name:", reflect.ValueOf(fn))
	if !(handlerType.Kind() == reflect.Func) {
		return fmt.Errorf("%s is not of type reflect.Func", reflect.TypeOf(fn).Kind())
	}
	eventName := handlerType.In(0).Name()

	var channelMap = bus.registry[eventName]
	fmt.Println("ChannelMap before unsubscribe: ", channelMap)
	for handler, channel := range channelMap {
		if reflect.ValueOf(fn) == handler {
			fmt.Println("Found handler to unregister")
			close(channel)
			fmt.Println(len(channelMap))
			delete(channelMap, handler)
			fmt.Println(len(channelMap))

		}

	}
	fmt.Println("ChannelMap after unsubscribe: ", channelMap)
	return nil
}

func Send(event interface{}) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	channelMap := bus.registry[reflect.TypeOf(event).Name()]

	for _, channel := range channelMap {
		channel <- event

	}

}

type eventHandler struct {
	callBack reflect.Value
}

func (h *eventHandler) OnEvent(channel <-chan interface{}) {
	for evt := range channel {
		h.callBack.Call([]reflect.Value{reflect.ValueOf(evt)})
	}
}
