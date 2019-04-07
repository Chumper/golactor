package dispatcher

type debugDispatcher struct{}

func (debugDispatcher) Dispatch(fn func()) {
	fn()
}

func (debugDispatcher) Throughput() int {
	return 1
}

// NewDebugDispatcher will return a dispatcher that will dispatch syncronized.
// No concurency is going on with this dispatcher, useful for debugging
func NewDebugDispatcher() Dispatcher {
	return &debugDispatcher{}
}

func NewDefaultDispatcher() Dispatcher {
	return &debugDispatcher{}
}
