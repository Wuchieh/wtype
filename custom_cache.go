package wtype

import (
	"time"
)

// CustomCache
//
//	Cache objects with fully customizable functions
type CustomCache[T any] struct {
	duration time.Duration

	setFunc           func(T, time.Duration)
	getFunc           func() T
	beforeSetDuration func(duration time.Duration) time.Duration
	resetTimer        func(time.Duration)
	stopTimer         func()
}

// NewCustomCache creates a CustomCache backed by the given functions.
func NewCustomCache[T any](
	duration time.Duration,
	setFunc func(T, time.Duration),
	getFunc func() T,
	beforeSetDuration func(duration time.Duration) time.Duration,
	resetTimer func(time.Duration),
	stopTimer func(),
) *CustomCache[T] {
	return &CustomCache[T]{
		duration:          duration,
		setFunc:           setFunc,
		getFunc:           getFunc,
		beforeSetDuration: beforeSetDuration,
		resetTimer:        resetTimer,
		stopTimer:         stopTimer,
	}
}

// Set stores the value using the configured set function.
func (c *CustomCache[T]) Set(t T) {
	if c.setFunc == nil {
		return
	}
	c.setFunc(t, c.duration)
}

// Get returns the value from the configured get function.
func (c *CustomCache[T]) Get() T {
	if c.getFunc == nil {
		return *new(T)
	}
	return c.getFunc()
}

// SetDuration sets the cache duration, applying the beforeSetDuration hook when
// provided.
func (c *CustomCache[T]) SetDuration(duration time.Duration) {
	if c.beforeSetDuration != nil {
		duration = c.beforeSetDuration(duration)
	}
	c.duration = duration
}

// ResetTimer resets the cache timer using the configured reset function.
func (c *CustomCache[T]) ResetTimer() {
	if c.resetTimer == nil {
		return
	}
	c.resetTimer(c.duration)
}

// StopTimer stops the cache timer using the configured stop function.
func (c *CustomCache[T]) StopTimer() {
	if c.stopTimer == nil {
		return
	}
	c.stopTimer()
}
