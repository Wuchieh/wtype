package wtype

import (
	"context"
	"sync"
)

// WaitGroup is a sync.WaitGroup with helpers for launching goroutines and
// waiting with a context.
type WaitGroup struct {
	sync.WaitGroup
}

// Go runs f in a new goroutine and tracks it with the wait group. A nil f is
// ignored.
func (w *WaitGroup) Go(f func()) {
	if f == nil {
		return
	}
	w.wg().Go(func() {
		f()
	})
}

// WaitContext waits for the wait group or until ctx is done, whichever happens
// first.
func (w *WaitGroup) WaitContext(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		defer close(done)
		w.wg().Wait()
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (w *WaitGroup) wg() *sync.WaitGroup {
	return &w.WaitGroup
}
