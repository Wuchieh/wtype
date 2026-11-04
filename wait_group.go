package wtype

import (
	"context"
	"sync"
)

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Go(f func()) {
	if f == nil {
		return
	}
	w.wg().Add(1)
	go func() {
		defer w.wg().Done()
		f()
	}()
}

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
