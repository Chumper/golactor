package process

type ActorProcessRef struct {
	p *ActorProcess
}

// Send will just delegate the message to the Send method of the ActorProcess
func (b *ActorProcessRef) Send(message interface{}) {
	b.p.Send(message)
}

func (b *ActorProcessRef) Poison() {
	b.p.Send(b.p.messages.Poison())
}

func (b *ActorProcessRef) Kill() {
	b.p.Send(b.p.messages.Kill())
}

// NewActorRef returns a reference to an actor
func newActorRef() *ActorProcessRef {
	return &ActorProcessRef{}
}
