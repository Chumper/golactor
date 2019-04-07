package channel

import (
	"errors"
	"time"
)

type Ref struct {
	channel chan interface{}
}

func (b *Ref) Send(message interface{}) {
	select {
	case b.channel <- message:
		return
	default:
		return
	}
}

func (b *Ref) Kill() {}

func (b *Ref) Poison() {}

func (b *Ref) Get() interface{} {
	select {
	case msg := <-b.channel:
		return msg
	case <-time.After(3 * time.Second):
		return errors.New("Timeout during message fetching")
	}
}

// NewRef will return a reference that can only store the given amount of messages.
// It can be used to receive a single response from an actor.
// Then it can be destroyed
func NewRef(size int) *Ref {
	return &Ref{
		channel: make(chan interface{}, size),
	}
}
