package mailbox

// A SystemMessage message is reserved for specific lifecycle messages used by the actor system
type SystemMessage interface {
	systemMessage()
}

// NewDefaultMailbox returns a mailbox that will prioritize system messages.
// For that it will maintain two queues, one for user messages and another for system messages.
// SystemMessages will be delivered first.
func NewDefaultMailbox() Mailbox {
	return &baseMailbox{
		userMailbox:   NewChannelQueue(2),
		systemMailbox: NewChannelQueue(2),
	}
}

type baseMailbox struct {
	userMailbox   Queue
	systemMailbox Queue
}

func (b *baseMailbox) Get() interface{} {
	// get system, if empty take user queue
	msg := b.systemMailbox.Pop()
	if msg == nil {
		return b.userMailbox.Pop()
	}
	return msg
}

func (b *baseMailbox) Store(msg interface{}) error {
	// determine if system or user message and post into queue
	switch msg := msg.(type) {
	case SystemMessage:
		return b.systemMailbox.Push(msg)
	default:
		return b.userMailbox.Push(msg)
	}
}
