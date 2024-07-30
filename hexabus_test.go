package hexades

import (
	"fmt"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestSubscribeUnsubscribeSend(t *testing.T) {
	data := MyData{}

	var handler = NewEventHandler(&data)
	Bus.Subscribe(handler)
	Bus.Send(MyEvent{"foo"})
	time.Sleep((100 * time.Millisecond))
	assert.Equal(t, "foo", data.val)

	Bus.Unsubscribe(handler)
	handler = NewEventHandler(&data)
	Bus.Send(MyEvent{"bar"})
	time.Sleep((100 * time.Millisecond))
	assert.Assert(t, !(data.val == "bar"))

	Bus.Subscribe(handler)
	Bus.Send(MyEvent{"bar"})
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, data.val, "bar")

}

type MyData struct{ val string }

func NewEventHandler(data *MyData) func(evt MyEvent) {
	return func(evt MyEvent) {
		data.val = evt.val
	}
}

func MyEventHandler(evt MyEvent) {
	fmt.Println("received in test: ", evt)

}

type MyEvent struct{ val string }
