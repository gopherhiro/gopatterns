package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	item := NewItem("Nike Shirt")

	observer1 := &Customer{}
	observer2 := &Trader{}

	item.register(observer1)
	item.register(observer2)

	item.haveNewEvent()

	item.deregister(observer1)
	item.haveNewEvent()
}
