package wtype

import (
	"context"
	"errors"
	"sync"
)

type ErrorGroup struct {
	wg         WaitGroup
	ctx        context.Context
	cancel     context.CancelCauseFunc
	errorAbort bool

	errorsMux sync.Mutex
	errors    []error
}

func (e *ErrorGroup) handleError(err error) {
	if err == nil {
		return
	}

	if e.errorAbort {
		e.cancel(err)
	} else {
		e.errorsMux.Lock()
		e.errors = append(e.errors, err)
		e.errorsMux.Unlock()
	}
}

func (e *ErrorGroup) Go(f func(context.Context) error) {
	if f == nil {
		return
	}

	e.wg.Go(func() {
		err := f(e.ctx)
		e.handleError(err)
	})
}

func (e *ErrorGroup) Wait() error {
	_ = e.wg.WaitContext(e.ctx)
	if cErr := context.Cause(e.ctx); cErr != nil &&
		!errors.Is(cErr, context.Canceled) {
		return cErr
	}

	if len(e.errors) > 0 {
		return errors.Join(e.errors...)
	}

	return nil
}

type ErrorGroupOption struct {
	ctx        context.Context
	errorAbort bool
}

func NewErrorGroup(options ...func(*ErrorGroupOption)) *ErrorGroup {
	option := UseOptions(ErrorGroupOption{}, options...)

	ctx := option.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancelCause(ctx)

	return &ErrorGroup{
		ctx:        ctx,
		cancel:     cancel,
		errorAbort: option.errorAbort,
	}
}
