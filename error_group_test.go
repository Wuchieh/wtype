package wtype_test

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

func TestErrorGroup_Success(t *testing.T) {
	const goroutines = 10

	tests := []struct {
		name    string
		options []func(*wtype.ErrorGroupOption)
	}{
		{name: "default"},
		{
			name:    "error abort",
			options: []func(*wtype.ErrorGroupOption){wtype.ErrorGroupErrorAbort(true)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := wtype.NewErrorGroup(tt.options...)
			var completed atomic.Int32

			for range goroutines {
				g.Go(func(ctx context.Context) error {
					if err := ctx.Err(); err != nil {
						return err
					}
					completed.Add(1)
					return nil
				})
			}

			if err := g.Wait(); err != nil {
				t.Fatalf("Wait() error = %v, want nil", err)
			}
			if got := completed.Load(); got != goroutines {
				t.Fatalf("completed goroutines = %d, want %d", got, goroutines)
			}
		})
	}
}

func TestErrorGroup_GoNil(t *testing.T) {
	g := wtype.NewErrorGroup()
	g.Go(nil)

	if err := g.Wait(); err != nil {
		t.Fatalf("Wait() error = %v, want nil", err)
	}
}

func TestErrorGroup_WaitWithoutGo(t *testing.T) {
	if err := wtype.NewErrorGroup().Wait(); err != nil {
		t.Fatalf("Wait() error = %v, want nil", err)
	}
}

func TestErrorGroup_CollectErrors(t *testing.T) {
	errA := errors.New("error a")
	errB := errors.New("error b")

	tests := []struct {
		name    string
		options []func(*wtype.ErrorGroupOption)
	}{
		{name: "default"},
		{
			name:    "explicitly disabled",
			options: []func(*wtype.ErrorGroupOption){wtype.ErrorGroupErrorAbort(false)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := wtype.NewErrorGroup(tt.options...)
			g.Go(func(context.Context) error { return errA })
			g.Go(func(context.Context) error { return errB })
			g.Go(func(context.Context) error { return nil })

			err := g.Wait()
			if !errors.Is(err, errA) {
				t.Errorf("Wait() error = %v, want to include %v", err, errA)
			}
			if !errors.Is(err, errB) {
				t.Errorf("Wait() error = %v, want to include %v", err, errB)
			}
		})
	}
}

func TestErrorGroup_CollectErrorsDoesNotCancel(t *testing.T) {
	wantErr := errors.New("boom")
	proceed := make(chan struct{})
	finished := make(chan struct{})

	g := wtype.NewErrorGroup()
	g.Go(func(ctx context.Context) error {
		<-proceed
		if err := ctx.Err(); err != nil {
			return err
		}
		close(finished)
		return nil
	})
	g.Go(func(context.Context) error {
		close(proceed)
		return wantErr
	})

	err := g.Wait()
	if !errors.Is(err, wantErr) {
		t.Fatalf("Wait() error = %v, want %v", err, wantErr)
	}

	select {
	case <-finished:
	default:
		t.Fatal("collect mode canceled the group after the first error")
	}
}

func TestErrorGroup_ErrorAbort(t *testing.T) {
	wantErr := errors.New("boom")
	started := make(chan struct{})
	finished := make(chan struct{})

	g := wtype.NewErrorGroup(wtype.ErrorGroupErrorAbort(true))
	g.Go(func(ctx context.Context) error {
		defer close(finished)
		close(started)
		<-ctx.Done()
		return ctx.Err()
	})

	<-started

	g.Go(func(context.Context) error {
		return wantErr
	})

	err := g.Wait()
	if !errors.Is(err, wantErr) {
		t.Fatalf("Wait() error = %v, want %v", err, wantErr)
	}
	if errors.Is(err, context.Canceled) {
		t.Fatal("Wait() returned context.Canceled instead of the abort cause")
	}

	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("goroutine did not observe the abort cancellation")
	}
}

func TestErrorGroup_WaitBlocks(t *testing.T) {
	g := wtype.NewErrorGroup()
	started := make(chan struct{})
	release := make(chan struct{})

	g.Go(func(context.Context) error {
		close(started)
		<-release
		return nil
	})
	<-started

	waited := make(chan error, 1)
	go func() {
		waited <- g.Wait()
	}()

	select {
	case err := <-waited:
		t.Fatalf("Wait() returned before goroutines finished: %v", err)
	case <-time.After(20 * time.Millisecond):
	}

	close(release)

	select {
	case err := <-waited:
		if err != nil {
			t.Fatalf("Wait() error = %v, want nil", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Wait() did not return after goroutines finished")
	}
}

func TestErrorGroup_WithContext(t *testing.T) {
	t.Run("nil uses background", func(t *testing.T) {
		g := wtype.NewErrorGroup(wtype.ErrorGroupWithContext(nil))
		g.Go(func(ctx context.Context) error {
			if ctx == nil {
				return errors.New("context is nil")
			}
			return ctx.Err()
		})

		if err := g.Wait(); err != nil {
			t.Fatalf("Wait() error = %v, want nil", err)
		}
	})

	t.Run("deadline exceeded", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		g := wtype.NewErrorGroup(wtype.ErrorGroupWithContext(ctx))
		<-ctx.Done()

		if err := g.Wait(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("Wait() error = %v, want %v", err, context.DeadlineExceeded)
		}
	})

	t.Run("parent canceled", func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		g := wtype.NewErrorGroup(wtype.ErrorGroupWithContext(parent))

		started := make(chan struct{})
		observed := make(chan struct{})
		g.Go(func(ctx context.Context) error {
			close(started)
			<-ctx.Done()
			close(observed)
			return nil
		})

		<-started
		cancel()
		<-observed

		// context.Canceled is not treated as an abort cause, and the goroutine
		// returned nil, so Wait reports no error.
		if err := g.Wait(); err != nil {
			t.Fatalf("Wait() error = %v, want nil", err)
		}
	})
}

func TestErrorGroup_ConcurrentErrors(t *testing.T) {
	const goroutines = 100

	g := wtype.NewErrorGroup()
	errs := make([]error, goroutines)
	var completed atomic.Int32

	for i := range goroutines {
		if i%2 == 0 {
			errs[i] = fmt.Errorf("error %d", i)
		}
		g.Go(func(context.Context) error {
			completed.Add(1)
			return errs[i]
		})
	}

	err := g.Wait()
	if got := completed.Load(); got != goroutines {
		t.Fatalf("completed goroutines = %d, want %d", got, goroutines)
	}
	for _, want := range errs {
		if want == nil {
			continue
		}
		if !errors.Is(err, want) {
			t.Errorf("Wait() error = %v, want to include %v", err, want)
		}
	}
}
