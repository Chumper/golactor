package mailbox

type ChannelQueue struct {
	channel chan interface{}
}

func (s *ChannelQueue) Push(msg interface{}) error {
	select {
	case s.channel <- msg:
		return nil
	default:
		return ErrrorQueueFull
	}
}
func (s *ChannelQueue) Pop() interface{} {
	select {
	case msg := <-s.channel:
		return msg
	default:
		return nil
	}
}

func NewChannelQueue(size int) Queue {
	return &ChannelQueue{
		channel: make(chan interface{}, size),
	}
}
