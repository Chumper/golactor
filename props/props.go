package props

import (
	"github.com/Chumper/golactor/actor"
)

// Option can be used to customize the props instance
type Option func(*BaseProp)

// FromFunc will create an actor prop from a receive function.
// It offers a quick way to create simple actors.
// Actors that require more complex interactions (e.g. persistence, mixins) are better created from a Producer.
func FromFunc(actorFunc actor.Func, opts ...Option) *BaseProp {
	b := &BaseProp{
		behavior: actorFunc,
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// FromProducer will create an actor prop from a producer function.
// It can be used to create more complex actors, as the actor itself can be stored in a separate file.
func FromProducer(producer actor.Producer, opts ...Option) *BaseProp {
	b := &BaseProp{
		behavior: producer(),
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}
