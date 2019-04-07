package actor

// A SystemMessage message is reserved for specific lifecycle messages used by the actor system
type SystemMessage interface {
	// This method is only there to provide correct type detection. It serves no other purpose
	SystemMessage()
}

// A Restarting message is sent to an actor when the actor is being restarted by the system due to a failure
type Restarting interface {
	// This method is only there to provide correct type detection. It serves no other purpose
	Restarting()
}

// A Stopping message is sent to an actor prior to the actor being stopped
type Stopping interface {
	// This method is only there to provide correct type detection. It serves no other purpose
	Stopping()
}

// A Stopped message is sent to the actor once it has been stopped. A stopped actor will receive no further messages
type Stopped interface {
	// This method is only there to provide correct type detection. It serves no other purpose
	Stopped()
}

// A Started message is sent to an actor once it has been started and ready to begin receiving messages.
type Started interface {
	// This method is only there to provide correct type detection. It serves no other purpose
	Started()
}

// A Watch message is sent to an actor when another actor is interested in the lifecycle events of the receiving actor.
// The message contains the actor.Ref of the actor that wants to be notified.
type Watch interface {
	Watcher() Ref
	// This method is only there to provide correct type detection. It serves no other purpose
	Watch()
}

// An Unwatch message is sent as soon as an actor is not interested anymore in the lifecycle of the receiving actor.
type Unwatch interface {
	Watcher() Ref
	// This method is only there to provide correct type detection. It serves no other purpose
	Unwatch()
}

// A Terminated message is sent to a watcher of an actor as soon as the watched actor dies.
type Terminated interface {
	Actor() Ref
	// This method is only there to provide correct type detection. It serves no other purpose
	Terminated()
}

// A Kill message is sent as a system message. It will indicate the actor to imediately shut down with noo further proccessing.
// All messages still in the queue will be send to the dead letter actor
type Kill interface {
	// This method is only there to provide correct type detection. It serves no other purpose
	Kill()
}

// A Poison message is sent as a user message, so it will retain the order of all messages and will indicate to the actor to shut down
type Poison interface {
	// This method is only there to provide correct type detection. It serves no other purpose
	Poison()
}

// MessageProvider provides system messages for further uses
type MessageProvider interface {
	Restarting() Restarting
	Stopping() Stopping
	Stopped() Stopped
	Started() Started
	Watch(Ref) Watch
	Unwatch(Ref) Unwatch
	Terminated(Ref) Terminated
	Kill() Kill
	Poison() Poison
}
