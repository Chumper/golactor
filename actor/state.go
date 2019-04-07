package actor

import (
	"github.com/Chumper/golactor/logging"
)

// State is the struct that every actor get passed into the receive method.
// It offers all methods that can be used to interact with the actor system.
type State interface {
	BasePart
	SpawnerPart
	// SenderPart
}

type BasePart interface {
	// Self returns the reference to the actor itself
	Self() Ref
	// Parent returns the reference to the parent actor
	Parent() Ref
	// Actor returns the actor associated with this state
	Actor() Actor
	// Children returns the actor references of all actor that this one has directly created
	Children() []Ref
}

type SpawnerPart interface {
	// Spawn will take the given Props and spawns an actor in the system.
	// It will return the actor.Ref of the created actor or an error if anything went wrong.
	// As it will always be called on an actor the parent is already determined.
	Spawn(Props) (Ref, error)
}

type LogginPart interface {
	// Log will return the logging instance to use. It will be configured with specific field for the actor
	Log() logging.Log
}
