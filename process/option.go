package process

import (
	"github.com/Chumper/golactor/actor"
)

func Parent(p actor.Ref) Option {
	return func(a *ActorProcess) {
		a.parent = p
	}
}
func Self(p actor.Ref) Option {
	return func(a *ActorProcess) {
		a.self = p
	}
}

func Messages(p actor.MessageProvider) Option {
	return func(a *ActorProcess) {
		a.messages = p
	}
}
