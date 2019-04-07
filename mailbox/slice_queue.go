package mailbox

type SliceQueue struct {
	queue []interface{}
}

func (s *SliceQueue) Push(msg interface{}) error {
	s.queue = append(s.queue, msg)
	return nil
}
func (s *SliceQueue) Pop() interface{} {
	if len(s.queue) > 0 {
		x, a := s.queue[0], s.queue[1:]
		s.queue = a
		return x
	}
	return nil
}
