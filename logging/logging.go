package logging

import (
	"go.uber.org/zap"
)

// Log is the interface all actors should use to log anything within the actor
// The interface makes sure that specific fields that defines the actor are added.
//
type Log interface {
	Info(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
}
