package deadletter

import (
	"errors"

	"github.com/Chumper/golactor/actor"
)

var (
	// ErrorMailboxTerminated will be thrown every time a message will stored
	ErrorMailboxTerminated = errors.New("Mailbox is terminated")
)

// Mailbox is a mailbox that redirects all messages to the deadletter actor
type Mailbox struct {
	deadLetter actor.Ref
}

// NewMailbox returns a new Mailbox that forwards all messages to the given actor
func NewMailbox(d actor.Ref) *Mailbox {
	return &Mailbox{
		deadLetter: d,
	}
}

// Get returns nil everytime as the messages are already forwarded
func (b *Mailbox) Get() interface{} {
	// return no message as all messages are already forwarded
	return nil
}

// Store will forward the given message to the dead letter actor
func (b *Mailbox) Store(msg interface{}) error {
	if b.deadLetter != nil {
		_, ok := msg.(actor.SystemMessage)
		if !ok {
			// no system message, so redirect
			b.deadLetter.Send(msg)
		}
	}
	return nil
}
