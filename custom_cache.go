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

func (c *CustomCache[T]) Set(t T) {
	if c.setFunc != nil {
		return
	}
	c.setFunc(t, c.duration)
}

func (c *CustomCache[T]) Get() T {
	if c.getFunc != nil {
		return *new(T)
	}
	return c.getFunc()
}

func (c *CustomCache[T]) SetDuration(duration time.Duration) {
	if c.beforeSetDuration != nil {
		duration = c.beforeSetDuration(duration)
	}
	c.duration = duration
}

func (c *CustomCache[T]) ResetTimer() {
	if c.resetTimer == nil {
		return
	}
	c.resetTimer(c.duration)
}

func (c *CustomCache[T]) StopTimer() {
	if c.stopTimer == nil {
		return
	}
	c.stopTimer()
}
