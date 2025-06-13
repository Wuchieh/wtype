package wtype

import (
	"time"
)

// CustomCache
//
//	Cache objects with fully customizable functions
type CustomCache[T any] struct {
	duration time.Duration

	setFunc     func(T)
	getFunc     func() T
	setDuration func(duration time.Duration)
	resetTimer  func()
	stopTimer   func()
}

func NewCustomCache[T any](
	setFunc func(T),
	getFunc func() T,
	setDuration func(duration time.Duration),
	resetTimer func(),
	stopTimer func()) *CustomCache[T] {
	return &CustomCache[T]{
		setFunc:     setFunc,
		getFunc:     getFunc,
		setDuration: setDuration,
		resetTimer:  resetTimer,
		stopTimer:   stopTimer,
	}
}

func (c *CustomCache[T]) Set(t T) {
	c.setFunc(t)
}

func (c *CustomCache[T]) Get() T {
	return c.getFunc()
}

func (c *CustomCache[T]) SetDuration(duration time.Duration) {
	c.setDuration(duration)
}

func (c *CustomCache[T]) ResetTimer() {
	c.resetTimer()
}

func (c *CustomCache[T]) StopTimer() {
	c.stopTimer()
}
