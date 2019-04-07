package actor

// Actor is the interface that defines the Receive method.
//
// Receive is sent messages to be processed from the mailbox associated with the instance of the actor
type Actor interface {
	// Receive is the core of each actor. All message handling happens here.
	// The actor can use the state to interact with the actor system.
	Receive(state State, message interface{})
}

// Func is a function that represents the base of an actor: the receive method
// It will be called to handle all incoming messages
type Func func(state State, message interface{})

// Receive calls f(c)
func (f Func) Receive(state State, message interface{}) {
	f(state, message)
}

// The Producer type is a function that creates a new actor
type Producer func() Actor

// ReceiverFunc is a function that will work on the actor state and the message.
// It acts as a middleware and must delegate the execution to the next Receiver
type ReceiverFunc func(state State, message interface{})

// Receiver is a chaing of Receivers that can be called
type Receiver func(next ReceiverFunc) ReceiverFunc

// Process is the heart of every actor. Here all the magic happens. This is the machine room of the actor.
type Process interface {
	Send(interface{})
}

// ProcessProducer is a function that can return a process based on the given props and Ref which acts as parent.
type ProcessProducer func(Props, Ref) Process

// Spawner defines a func that is able to spawn new actors and returns the Ref to the actor or an error
type Spawner func(props Props, parent Ref) (Ref, error)
