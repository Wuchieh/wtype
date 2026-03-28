package wtype_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

func TestWaitGroup_WaitContext(t *testing.T) {
	var wg wtype.WaitGroup

	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	wg.Go(func() {
		time.Sleep(1 * time.Second)
	})

	if err := wg.WaitContext(ctx); err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Error("WaitContext error:", err)
		}
	}

	since := time.Since(now)
	if since > time.Second {
		t.Error("since should be within 1s")
	}
}
