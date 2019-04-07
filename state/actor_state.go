package state

import (
	"github.com/Chumper/golactor/actor"
)

// ActorState is the state that can be used in an actor to interact with the system
type ActorState struct {
	MessagesFunc  func() actor.MessageProvider
	Spawner       actor.Spawner
	BehaviourFunc func() actor.Actor
	ParentFunc    func() actor.Ref
	SelfFunc      func() actor.Ref
	ChildrenFunc  func() []actor.Ref
}

// New returns a new state that can be used in an actor
func New(opts ...Option) (*ActorState, error) {
	// func New(behaviour actor.Actor, parent actor.Ref, self actor.Ref, spawner actor.Spawner, messages actor.MessageProvider) *ActorState {
	s := &ActorState{}

	// apply options
	for _, opt := range opts {
		opt(s)
	}

	err := s.verify()

	return s, err
}

func (a *ActorState) verify() error {
	return nil
}

func (a *ActorState) Children() []actor.Ref {
	return a.ChildrenFunc()
}

func (a *ActorState) Self() actor.Ref {
	return a.SelfFunc()
}
func (a *ActorState) Parent() actor.Ref {
	return a.ParentFunc()
}
func (a *ActorState) Actor() actor.Actor {
	// returns the actor
	return a.BehaviourFunc()
}
func (a *ActorState) Spawn(props actor.Props) (actor.Ref, error) {
	ref, err := a.Spawner(props, a.Self())
	if err != nil {
		return nil, err
	}
	// watch child
	ref.Send(a.MessagesFunc().Watch(a.Self()))
	return ref, err
}
