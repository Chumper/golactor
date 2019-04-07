package system

import (
	"fmt"

	"github.com/Chumper/golactor/actor"
	"github.com/Chumper/golactor/props"
)

// guardianActor is the root actor for each actor system.
// If the guardianActor fails, everything fails...
type guardianActor struct {
}

// spawnChild is a simple message that will spawn a new child in the context of an actor
type spawnChild struct {
	RespondTo actor.Ref
	Props     actor.Props
}

// spawnChildResponse is the response that will be send on the request to spwan a child
type spawnChildResponse struct {
	ActorRef actor.Ref
}

// Receive can handle no message expect adding new actors as child
func (g *guardianActor) Receive(state actor.State, message interface{}) {
	switch msg := message.(type) {
	case *spawnChild:
		childRef, err := state.Spawn(msg.Props)
		if err != nil {
			// SHUTDOWN!!! errors are not allowed here
			panic("Spawning child failed, that should not happen in the Guardian actor, shuting down...")
		}
		// reply to sender
		msg.RespondTo.Send(&spawnChildResponse{childRef})
	}

}

// NewguardianActor creates a new guardianActor and returns the Actor that can be used in Props
func newGuardianActor(b *BaseSystem) *props.BaseProp {
	return props.FromProducer(func() actor.Actor { return &guardianActor{} },
		props.MessageProvider(b.messages),
		props.Dispatcher(b.dispatcher),
	)
}

// systemActor is the root actor for each actor system.
// If the systemActor fails, it will escalate to the guardian actor
type systemActor struct{}

// Receive can handle no message expect adding new actors as child
func (g *systemActor) Receive(state actor.State, message interface{}) {
	switch msg := message.(type) {
	case *spawnChild:
		childRef, err := state.Spawn(msg.Props)
		if err != nil {
			// send err to the actor
			msg.RespondTo.Send(err)
		}
		// reply to sender
		msg.RespondTo.Send(&spawnChildResponse{childRef})
	}

}

// NewsystemActor creates a new systemActor and returns the Actor that can be used in Props
func newSystemActor(b *BaseSystem) *props.BaseProp {
	return props.FromProducer(func() actor.Actor { return &systemActor{} },
		props.MessageProvider(b.messages),
		props.DeadLetter(b.deadLetterRef),
	)
}

// userActor is the root actor for each actor system.
// If the userActor fails, it will escalate to the guardian actor
type userActor struct {
}

// Receive can handle no message expect adding new actors as child
func (g *userActor) Receive(state actor.State, message interface{}) {
	switch msg := message.(type) {
	// Request to spawn a child actor initiated by the user
	case *spawnChild:
		childRef, err := state.Spawn(msg.Props)
		if err != nil {
			// send err to the actor
			msg.RespondTo.Send(err)
		}
		// reply to sender
		msg.RespondTo.Send(&spawnChildResponse{childRef})
	case actor.Terminated:
		fmt.Println("Child terminated, left:", len(state.Children()))
	}
}

// NewuserActor creates a new userActor and returns the Actor that can be used in Props
func newUserActor(b *BaseSystem) *props.BaseProp {
	return props.FromProducer(func() actor.Actor { return &userActor{} },
		props.MessageProvider(b.messages),
		props.DeadLetter(b.deadLetterRef),
	)
}

// ---------------------

// deadLetterActor is the actor that will receive all messages that can not be routed to a given reference
type deadLetterActor struct {
	// potential subscribers
}

// Receive will redistribute messages to subscribers and log them
func (g *deadLetterActor) Receive(state actor.State, message interface{}) {
	fmt.Println("[DEADLETTER] Received a dead letter message:", message)
}

// newDeadLetterActor creates a new deadLetter and returns the Actor that can be used in Props
func newDeadLetterActor(b *BaseSystem) *props.BaseProp {
	return props.FromProducer(
		func() actor.Actor { return &deadLetterActor{} },
		props.MessageProvider(b.messages),
	)
}
