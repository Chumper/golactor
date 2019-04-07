package mailbox

import (
	"errors"
)

var (
	// ErrrorQueueFull will be thrown if the queue is full
	ErrrorQueueFull = errors.New("Queue is full")
)

type Queue interface {
	Push(interface{}) error
	Pop() interface{}
}
