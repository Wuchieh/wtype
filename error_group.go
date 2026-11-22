package wtype

import (
	"context"
	"errors"
	"sync"
)

// ErrorGroup runs a group of goroutines and collects the errors they return.
//
// When errorAbort is disabled (the default), every non-nil error is collected
// and Wait joins them into a single error. When errorAbort is enabled, the
// first non-nil error cancels the group's context with that error as the cause,
// and Wait returns it.
type ErrorGroup struct {
	wg         WaitGroup
	ctx        context.Context
	cancel     context.CancelCauseFunc
	errorAbort bool

	errorsMux sync.Mutex
	errors    []error
}

// handleError records err according to the group's error mode. A nil error is
// ignored.
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

// Go runs f in a new goroutine and passes the group's context to it. The error
// returned by f is handled according to the group's error mode. A nil f is
// ignored.
func (e *ErrorGroup) Go(f func(context.Context) error) {
	if f == nil {
		return
	}

	e.wg.Go(func() {
		err := f(e.ctx)
		e.handleError(err)
	})
}

// Wait blocks until every goroutine started by Go has finished or the group's
// context is canceled, whichever happens first.
//
// If the group's context was canceled with a cause other than
// context.Canceled, that cause is returned. Otherwise, all errors collected
// while errorAbort was disabled are joined into a single error. Wait returns nil
// when no error occurred.
func (e *ErrorGroup) Wait() error {
	// WaitContext only reports cancellation; the cause is checked below.
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

// ErrorGroupOption holds the configuration applied by NewErrorGroup. A zero
// value uses context.Background and collects all errors instead of aborting.
type ErrorGroupOption struct {
	ctx        context.Context
	errorAbort bool
}

// NewErrorGroup creates an ErrorGroup. The given options are applied in order
// to an ErrorGroupOption before the group is created.
//
// The group's context is derived from the option's context, which defaults to
// context.Background when nil, so canceling the parent context also cancels the
// group.
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
