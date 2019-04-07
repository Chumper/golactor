package system

import (
	"time"

	"github.com/Chumper/golactor/actor"
	"github.com/Chumper/golactor/dispatcher"
	"go.uber.org/zap"
)

// Logger provides a custom logging provider
func Logger(l *zap.Logger) Option {
	return func(s *BaseSystem) {
		s.logger = l
	}
}

// Dispatcher provides a custom dispatcher
func Dispatcher(d dispatcher.Dispatcher) Option {
	return func(s *BaseSystem) {
		s.dispatcher = d
	}
}

// MessageProvider can be used to provide a custom MessageProvider
func MessageProvider(m actor.MessageProvider) Option {
	return func(s *BaseSystem) {
		s.messages = m
	}
}

func Timeout(m time.Duration) Option {
	return func(s *BaseSystem) {
		s.timeout = m
	}
}
