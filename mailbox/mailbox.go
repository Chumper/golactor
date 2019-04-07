package mailbox

// Mailbox is an interface that is used to enqueue messages to the mailbox
// A mailbox must be able to handle concurrent stores.
// Reads will always be syncronized.
type Mailbox interface {
	// Store will store the given message in the mailbox.
	// Will be used concurrently and must be safe to be used from multiple goroutines
	Store(message interface{}) error
	// Will get a message from the mailbox. Order is defined by the mailbox itself.
	Get() (message interface{})
}

// Producer is a function which creates a new mailbox
type Producer func() Mailbox
