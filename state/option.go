package state

import (
	"github.com/Chumper/golactor/actor"
)

// Option can be used to customize the props instance
type Option func(*ActorState)

func Parent(r actor.Ref) Option {
	return func(a *ActorState) {
		a.ParentFunc = func() actor.Ref { return r }
	}
}

func Children(f func() []actor.Ref) Option {
	return func(a *ActorState) {
		a.ChildrenFunc = f
	}
}

func Self(r actor.Ref) Option {
	return func(a *ActorState) {
		a.SelfFunc = func() actor.Ref { return r }
	}
}

func Actor(r actor.Actor) Option {
	return func(a *ActorState) {
		a.BehaviourFunc = func() actor.Actor { return r }
	}
}

func Spawner(s actor.Spawner) Option {
	return func(a *ActorState) {
		a.Spawner = s
	}
}

func Messages(m actor.MessageProvider) Option {
	return func(a *ActorState) {
		a.MessagesFunc = func() actor.MessageProvider { return m }
	}
}
