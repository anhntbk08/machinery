package tasks

import (
	"fmt"
	"time"
)

// ErrRetryTaskLater ...
type ErrRetryTaskLater struct {
	name, msg string
	retryIn   time.Duration
}

// RetryIn returns time.Duration from now when task should be retried
func (e ErrRetryTaskLater) RetryIn() time.Duration {
	return e.retryIn
}

// Error implements the error interface
func (e ErrRetryTaskLater) Error() string {
	return fmt.Sprintf("Task error: %s Will retry in: %s", e.msg, e.retryIn)
}

// NewErrRetryTaskLater returns new ErrRetryTaskLater instance
func NewErrRetryTaskLater(msg string, retryIn time.Duration) ErrRetryTaskLater {
	return ErrRetryTaskLater{msg: msg, retryIn: retryIn}
}

// Retriable is interface that retriable errors should implement
type Retriable interface {
	RetryIn() time.Duration
}

// ErrShouldStopProcessing ...
type ErrShouldStopProcessing struct {
	name, msg string
}

// Error implements the error interface
func (e ErrShouldStopProcessing) Error() string {
	return fmt.Sprintf("Task error: %s Will be stopped in: %s", e.msg)
}

func NewErrShouldStopProcessing(msg string) ErrShouldStopProcessing {
	return ErrShouldStopProcessing{msg: msg}
}
