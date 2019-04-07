package props

import (
	"github.com/Chumper/golactor/actor"
	"github.com/Chumper/golactor/dispatcher"
	"github.com/Chumper/golactor/mailbox"
)

// Dispatcher provides a custom dispatcher
func Dispatcher(d dispatcher.Dispatcher) Option {
	return func(s *BaseProp) {
		s.dispatcher = d
	}
}

// Mailbox provides a custom Mailbox
func Mailbox(m mailbox.Mailbox) Option {
	return func(s *BaseProp) {
		s.mailbox = m
	}
}

// Receiver provides the Props with custom receiver middleware
func Receiver(receivers ...actor.Receiver) Option {
	return func(s *BaseProp) {
		s.receivers = receivers
	}
}

// MessageProvider provides a custom system message generator
func MessageProvider(m actor.MessageProvider) Option {
	return func(s *BaseProp) {
		s.messages = m
	}
}

func DeadLetter(m actor.Ref) Option {
	return func(s *BaseProp) {
		s.deadLetter = m
	}
}
