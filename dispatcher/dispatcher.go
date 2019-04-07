package dispatcher

// Dispatcher is the interface to dispatch all messages.
// They can be choosen on actor level and determine how messages are dispatched.
type Dispatcher interface {
	// Dispatch will execute the given function as defined by dispatcher implementation
	Dispatch(fn func())
	// Throughput defines the max amount of functions to be processed by the dispatcher before it allows other goroutines to run
	Throughput() int
}
