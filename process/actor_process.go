package process

import (
	"fmt"
	"sync/atomic"

	"github.com/Chumper/golactor/deadletter"

	"github.com/Chumper/golactor/state"

	"github.com/Chumper/golactor/actor"
	"github.com/Chumper/golactor/dispatcher"

	"github.com/Chumper/golactor/mailbox"
)

const (
	// mailIdle indicates that the actor process is currently not running
	mailboxIdle = 0
	// mailRunning indicates that the actor process is currently running and processing messages
	mailboxRunning = 1
	// mailboxTerminated indicates that the mailbox and therefore the actor has been terminated
	mailboxTerminated = 2
)

// ActorProcess is where all the dirty works happens
type ActorProcess struct {
	// status identifies if the process is running or idle
	status int32
	// messageCounter defines the amount on messages still to be processed
	messageCounter int32
	// mailbox will contain all messages that still need to be processed
	mailbox mailbox.Mailbox
	// dispatcher defines the dispatcher to use for processing the messages
	dispatcher dispatcher.Dispatcher
	// state is the interface that offers any interaction with the system
	state *state.ActorState
	// receiver represents the middle ware that will be called before the actual state of the actor will be called
	receiver actor.ReceiverFunc
	// terminated determines if the process is terminated or not
	terminated bool
	// deadLetter is the reference to the dead letter actor where all messages are sent to after termination
	deadLetter actor.Ref
	// parent is the reference to the parent actor if any
	parent actor.Ref
	// self is the reference to the actor itself
	self actor.Ref
	// children contains a list of children references
	children []actor.Ref
	// watcher contains a list of watcher references
	watcher []actor.Ref
	// messages is the message provider providing system messages
	messages actor.MessageProvider
}

func New(props actor.Props, opts ...Option) *ActorProcess {
	p := &ActorProcess{
		dispatcher: props.Dispatcher(),
		mailbox:    props.Mailbox(),
		deadLetter: props.DeadLetter(),
	}

	// apply options
	for _, opt := range opts {
		opt(p)
	}

	// apply defaults
	p.applyDefaults()

	return p
}

func (a *ActorProcess) applyDefaults() {
	if a.mailbox == nil {
		a.mailbox = mailbox.NewDefaultMailbox()
	}
	if a.dispatcher == nil {
		a.dispatcher = dispatcher.NewDefaultDispatcher()
	}
}

// Send will send the given message to the actor behind the process
func (a *ActorProcess) Send(message interface{}) {
	// run process with the dispatcher
	a.dispatcher.Dispatch(func() {
		// store the message
		if err := a.mailbox.Store(message); err == nil {
			// increase message counter
			atomic.AddInt32(&a.messageCounter, 1)
			// run the main loop if needed
			a.run()
		}
	})
}

func (a *ActorProcess) run() {
	// only run when we can get the permission to run
	started := atomic.CompareAndSwapInt32(&a.status, mailboxIdle, mailboxRunning)
	if started {
		// get message from mailbox
		for msg := a.mailbox.Get(); msg != nil; msg = a.mailbox.Get() {
			// decrease message count
			atomic.AddInt32(&a.messageCounter, -1)
			// invoke message
			a.invoke(msg)
		}
		// no more messages, set to idle if running
		// if the actor has been terminated, then the mailbox is also terminated
		idle := atomic.CompareAndSwapInt32(&a.status, mailboxRunning, mailboxIdle)
		if idle {
			// check if messages have been added
			if atomic.LoadInt32(&a.messageCounter) > 0 {
				//rerun the proccess
				a.run()
			}
		}
	}
}

func (a *ActorProcess) watch(ref actor.Ref) {
	// add watcher
	a.watcher = append(a.watcher, ref)
}

func (a *ActorProcess) unwatch(ref actor.Ref) {
	// remove watcher
	for i, r := range a.watcher {
		if r == ref {
			a.watcher = append(a.watcher[:i], a.watcher[i+1:]...)
			return
		}
	}
}

func (a *ActorProcess) terminate() {
	// terminate right now!!!
	fmt.Println("TERMINATING ACTOR!!!")
	a.terminated = true

	// invoke stopping message
	a.invoke(a.messages.Stopping())

	// redirect all messages to dead letter aka. replace mailbox

	m := a.mailbox
	// create a new mailbox so that messages will be redirected
	a.mailbox = deadletter.NewMailbox(a.deadLetter)
	if a.deadLetter != nil {
		// drain the current mailbox
		for msg := m.Get(); msg != nil; msg = m.Get() {
			// if this is a system message, then do not forward it
			_, ok := msg.(actor.SystemMessage)
			if !ok {
				// no system message, so forward
				a.deadLetter.Send(msg)
			}
		}
	}
	// send stopped message as there are no more actions happening after this
	a.invoke(a.messages.Stopped())

	// set mailbox to terminated
	atomic.StoreInt32(&a.status, mailboxTerminated)

	// notify watchers
	for _, w := range a.watcher {
		// send a Terminated message
		w.Send(a.messages.Terminated(a.self))
	}
}

func (a *ActorProcess) invoke(msg interface{}) {
	// send through middleware
	a.receiver(a.state, msg)
	// send to actor
	a.state.Actor().Receive(a.state, msg)
}

func (a *ActorProcess) handleSystemMessages() actor.ReceiverFunc {
	return func(state actor.State, message interface{}) {
		switch msg := message.(type) {
		case actor.Kill:
		case actor.Poison:
			// kill actor
			a.terminate()
			break
		// handle watch
		case actor.Watch:
			// add watcher
			a.watch(msg.Watcher())
			break
		// handle unwatch
		case actor.Unwatch:
			a.unwatch(msg.Watcher())
			break
		}
	}
}

// will create a receiver chain that can interact on messages. It will add the system message handling at the end
func (a *ActorProcess) makeReceiverChain(receiver []actor.Receiver) actor.ReceiverFunc {
	if len(receiver) == 0 {
		return a.handleSystemMessages()
	}

	h := receiver[len(receiver)-1](a.handleSystemMessages())
	for i := len(receiver) - 2; i >= 0; i-- {
		h = receiver[i](h)
	}
	return h
}
