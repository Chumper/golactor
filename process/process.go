package process

import (
	"github.com/Chumper/golactor/actor"
	"github.com/Chumper/golactor/state"
)

// Option defines the options available on the Process
type Option func(*ActorProcess)

func InitialSpawn(props actor.Props, parent actor.Ref, deadLetter actor.Ref, messages actor.MessageProvider) (actor.Ref, error) {
	// create actor ref first
	ref := newActorRef()

	// create process
	process := New(props,
		Parent(parent),
		Self(ref),
		Messages(messages),
	)

	// set dead letter reference
	if process.deadLetter == nil {
		process.deadLetter = deadLetter
	}

	// set receiver chain
	process.receiver = process.makeReceiverChain(props.Receiver())

	// assign process
	ref.p = process

	// create state
	s, err := state.New(
		state.Actor(props.Actor()),
		state.Messages(messages),
		state.Parent(parent),
		state.Self(ref),
		state.Spawner(process.Spawn),
		state.Children(func() []actor.Ref { return process.children }),
	)

	process.state = s

	return ref, err
}

// Spawn is a shortcut to create actors
func (a *ActorProcess) Spawn(props actor.Props, parent actor.Ref) (actor.Ref, error) {
	return InitialSpawn(props, parent, a.deadLetter, a.messages)
}

func applyDefaults(a *ActorProcess) {
}
