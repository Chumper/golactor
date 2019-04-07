package actor

// Ref is a reference to the actor behind the reference.
// It can be used to send messages to the actor process
// It does not offer any public methods yet
// With an interface we can support multiple Reference implementations. e.g. Protobuf,bson,flatbuffers
type Ref interface {
	// send is the internal method that all process implement to support sending messages to an actor.
	// there can be different processes for system, local, remote, deadletter or cluster references
	Send(interface{})
	// Poision will terminate the actor. Previous messages will still be handled
	// Internally this will just send a poison message
	Poison()
	// Will kill the actor immediately. Kill is sending a system message that will be handled
	// before any user message
	Kill()
	// Path() Paths
}

// Path represents the path of an actor including host and the internal path
// type Path interface {
// 	Host() string
// 	Path() string
// }
