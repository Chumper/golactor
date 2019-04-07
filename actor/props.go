package actor

import (
	"github.com/Chumper/golactor/dispatcher"
	"github.com/Chumper/golactor/mailbox"
)

// Props represents configuration to define how an actor should be created
type Props interface {
	// Actor returns the behaviour of the actor to be created
	Actor() Actor
	// Mailbox is the mailbox implementation to use for the actor
	Mailbox() mailbox.Mailbox
	// Dispatcher is the dispatcher to use for the actor
	Dispatcher() dispatcher.Dispatcher
	// Receiver return the receiver middleware that will interact on incoming messages before the actor gets the message
	Receiver() []Receiver
	// Returns a pointer to a message provider
	Messages() MessageProvider
	// Returns the reference to the deadletter actor
	DeadLetter() Ref
}
