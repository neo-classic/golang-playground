package events

import (
	"fmt"
	"testing"
)

type TestEvent struct {
	userID int
}

func (e TestEvent) Name() string {
	return "event.test"
}

func (e TestEvent) UserID() int {
	return e.userID
}

type TestEventHandler struct {
}

func (h TestEventHandler) Notify(event Event) {
	fmt.Printf("TestEventHandler %+v\n", event)
}

func TestEvents(t *testing.T) {
	publisher := NewEventPublisher()
	var eventHandler TestEventHandler
	e := TestEvent{}

	publisher.Subscribe(eventHandler, e)

	publisher.Notify(TestEvent{
		userID: 1,
	})
}
