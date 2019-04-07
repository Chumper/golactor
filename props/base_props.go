package props

import (
	"github.com/Chumper/golactor/actor"
	"github.com/Chumper/golactor/dispatcher"
	"github.com/Chumper/golactor/mailbox"
)

// BaseProp defines the base implementation of a Props.
// Has an actorFunc that defines the behavior of the actor
type BaseProp struct {
	behavior   actor.Actor
	dispatcher dispatcher.Dispatcher
	mailbox    mailbox.Mailbox
	receivers  []actor.Receiver
	messages   actor.MessageProvider
	deadLetter actor.Ref
}

func (b *BaseProp) Actor() actor.Actor {
	return b.behavior
}

func (b *BaseProp) Messages() actor.MessageProvider {
	return b.messages
}

func (b *BaseProp) Mailbox() mailbox.Mailbox {
	return b.mailbox
}

func (b *BaseProp) Dispatcher() dispatcher.Dispatcher {
	return b.dispatcher
}

func (b *BaseProp) Receiver() []actor.Receiver {
	return b.receivers
}

func (b *BaseProp) DeadLetter() actor.Ref {
	return b.deadLetter
}
