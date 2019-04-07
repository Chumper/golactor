package system

import (
	"errors"
	"time"

	"github.com/Chumper/golactor/messages"

	"github.com/Chumper/golactor/process"

	"github.com/Chumper/golactor/actor"
	"github.com/Chumper/golactor/channel"
	"github.com/Chumper/golactor/dispatcher"

	"go.uber.org/zap"
)

// BaseSystem is the base implementation of the System interface
// It spawns three actors on start. The Guardian actor, the user actor and the system actor
type BaseSystem struct {
	// Represents the "root" actor, all other actors are children of the guardian actor
	guardianRef actor.Ref
	// All user actors are children of the user actor
	userRef actor.Ref
	// All system internal actors are children of the system actor, e.g. deadletter actors
	systemRef actor.Ref
	// All messages that have no receiver are routed to the dead letter actor
	deadLetterRef actor.Ref
	// Name is the name of the sytem
	name string
	// Logger is the logger instance to use for all loggers
	logger *zap.Logger
	// Serializer holds the reference to all system messages as well as to the serialize methods used in remoting
	messages actor.MessageProvider
	// Dispatcher is the custom dispatcher used for all spawned actors if none is given
	dispatcher dispatcher.Dispatcher
	// spawner is the default method used to spawn
	spawner actor.Spawner
	// default timeout for all actions
	timeout time.Duration
}

// Option defines the options available on the system
type Option func(*BaseSystem)

// New will return a new actor system an marks the entry point for all actions with it.
func New(name string, opts ...Option) (*BaseSystem, error) {
	bs := &BaseSystem{
		name: name,
	}

	// apply options
	for _, opt := range opts {
		opt(bs)
	}

	// provide defaults if not given
	bs.provideDefaults()

	// initialize system
	err := bs.init()

	if err != nil {
		return nil, err
	}

	return bs, nil
}

// provides the defaults if none are given
func (b *BaseSystem) provideDefaults() {
	// provide default dispatcher
	if b.dispatcher == nil {
		b.dispatcher = dispatcher.NewGoRoutineDispatcher(10)
	}
	// provide default message provider
	if b.messages == nil {
		b.messages = messages.NewGoMessageProvider()
	}
	// provide default logger
	// provide default timeout
	if b.timeout == 0 {
		b.timeout = 5 * time.Second
	}

}

//
func (b *BaseSystem) init() error {
	if err := b.spawnGuardian(); err != nil {
		return err
	}
	if err := b.spawnDeadLetter(); err != nil {
		return err
	}
	if err := b.spawnSystem(); err != nil {
		return err
	}
	if err := b.spawnUser(); err != nil {
		return err
	}
	b.logger.Debug("Started actor system", zap.String("name", b.name))
	return nil
}

//
func (b *BaseSystem) Spawn(props actor.Props) (actor.Ref, error) {
	// spawn a new actor in the system under the user actor
	// just send a message to the user actor to spawn a child
	responseRef := channel.NewRef(1)
	b.userRef.Send(&spawnChild{Props: props, RespondTo: responseRef})
	switch msg := responseRef.Get().(type) {
	case *spawnChildResponse:
		return msg.ActorRef, nil
	case error:
		return nil, msg
	default:
		return nil, errors.New("Could not spawn the actor")
	}
}

// Send will just send the message to actor behind the reference
func (b *BaseSystem) Send(target actor.Ref, message interface{}) {
	target.Send(message)
}

func (b *BaseSystem) spawnGuardian() error {
	guardianRef, err := process.InitialSpawn(newSystemActor(b), nil, nil, b.messages)
	if err != nil {
		return err
	}
	b.guardianRef = guardianRef
	return nil
}

func (b *BaseSystem) spawnDeadLetter() error {

	// create response ref
	responseRef := channel.NewRef(1)

	// spawn dead letter actor
	b.Send(b.guardianRef, &spawnChild{responseRef, newDeadLetterActor(b)})
	// get ref from user actor
	switch msg := responseRef.Get().(type) {
	case *spawnChildResponse:
		b.deadLetterRef = msg.ActorRef
	case error:
		return msg
	}
	return nil
}

func (b *BaseSystem) spawnSystem() error {
	// create response ref
	responseRef := channel.NewRef(1)

	// spawn system actor
	b.Send(b.guardianRef, &spawnChild{responseRef, newSystemActor(b)})
	// get ref from user actor
	switch msg := responseRef.Get().(type) {
	case *spawnChildResponse:
		b.systemRef = msg.ActorRef
	case error:
		return msg
	}
	return nil
}

func (b *BaseSystem) spawnUser() error {
	// create response ref
	responseRef := channel.NewRef(1)

	// spawn user actor
	b.Send(b.guardianRef, &spawnChild{responseRef, newUserActor(b)})
	// get ref from user actor
	switch msg := responseRef.Get().(type) {
	case *spawnChildResponse:
		b.userRef = msg.ActorRef
	case error:
		return msg
	}
	return nil
}
