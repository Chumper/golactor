package system

import (
	"github.com/Chumper/golactor/actor"
)

// System is the root of all actors.
// Behind the system there is basically just another actor that marks the root for all user actors.
// The System interface enables the user to interact with system on behalf of the user actor.
type System interface {
	actor.SpawnerPart
	// actor.SenderPart
	// messages() messages.Messages
}

// // SystemBuilder is a builder that can be used to build an actor system.
// type SystemBuilder struct {
// 	// Name defines the name of the system itself, it will be part of the actor paths
// 	Name string
// 	// Dispatcher represents the default dispatcher used for all actors.
// 	// This can always be overwritten.
// 	Dispatcher dispatcher.Dispatcher
// 	// MaxCores defines the maximum amount of cores that the system will use
// 	MaxCores uint
// 	// Defines the logger to use for all actor related logging
// 	Logger *zap.Logger
// 	// Serializer is a factory to provide all system messages
// 	MessageProvider MessageProvider
// }
