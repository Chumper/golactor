package dispatcher

type goroutineDispatcher struct {
	throughput int
}

func (d *goroutineDispatcher) Throughput() int {
	return d.throughput
}

func (goroutineDispatcher) Dispatch(fn func()) {
	go fn()
}

// NewGoRoutineDispatcher will return a dispatcher that will dispatch with go functions.
func NewGoRoutineDispatcher(throughput int) Dispatcher {
	return &goroutineDispatcher{
		throughput: throughput,
	}
}
