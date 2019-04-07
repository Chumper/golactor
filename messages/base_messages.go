package messages

import (
	"github.com/Chumper/golactor/actor"
)

// Messages

// A Restarting message is sent to an actor when the actor is being restarted by the system due to a failure
type restarting struct{}

func (r *restarting) SystemMessage() {}
func (r *restarting) Restarting()    {}

// A Stopping message is sent to an actor prior to the actor being stopped
type stopping struct{}

func (s *stopping) SystemMessage() {}
func (s *stopping) Stopping()      {}

// A Stopped message is sent to the actor once it has been stopped. A stopped actor will receive no further messages
type stopped struct{}

func (s *stopped) SystemMessage() {}
func (s *stopped) Stopped()       {}

// A Started message is sent to an actor once it has been started and ready to begin receiving messages.
type started struct{}

func (s *started) SystemMessage() {}
func (s *started) Started()       {}

// Watch
type watch struct {
	_watcher actor.Ref
}

func (w *watch) SystemMessage() {}
func (w *watch) Watch()         {}
func (w *watch) Watcher() actor.Ref {
	return w._watcher
}

// Unwatch
type unwatch struct {
	_watcher actor.Ref
}

func (w *unwatch) SystemMessage() {}
func (w *unwatch) Unwatch()       {}
func (w *unwatch) Watcher() actor.Ref {
	return w._watcher
}

// Terminated
type terminated struct {
	_who actor.Ref
}

func (w *terminated) SystemMessage() {}
func (w *terminated) Terminated()    {}
func (w *terminated) Actor() actor.Ref {
	return w._who
}

// Poison
type poison struct{}

func (s *poison) Poison() {}

// Kill
type kill struct{}

func (s *kill) SystemMessage() {}
func (s *kill) Kill()          {}

type GoMessageProvider struct{}

func (g *GoMessageProvider) Started() actor.Started {
	return &started{}
}

func (g *GoMessageProvider) Restarting() actor.Restarting {
	return &restarting{}
}

func (g *GoMessageProvider) Stopping() actor.Stopping {
	return &stopping{}
}

func (g *GoMessageProvider) Stopped() actor.Stopped {
	return &stopped{}
}

func (g *GoMessageProvider) Poison() actor.Poison {
	return &poison{}
}

func (g *GoMessageProvider) Kill() actor.Kill {
	return &kill{}
}

func (g *GoMessageProvider) Watch(ref actor.Ref) actor.Watch {
	return &watch{ref}
}

func (g *GoMessageProvider) Unwatch(ref actor.Ref) actor.Unwatch {
	return &unwatch{ref}
}

func (g *GoMessageProvider) Terminated(who actor.Ref) actor.Terminated {
	return &terminated{who}
}

func NewGoMessageProvider() *GoMessageProvider {
	return &GoMessageProvider{}
}
