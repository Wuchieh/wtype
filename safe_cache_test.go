package wtype_test

import (
	"sync"
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

func TestSafeCache(t *testing.T) {
	s := wtype.NewSafeCache(time.Second, 0)
	s2 := 0
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(index int) {
			defer wg.Done()
			s2++
			s.Use(func(data int) int {
				return data + 1
			})
		}(i)
	}
	wg.Wait()
	if s.Get() != 1000 {
		t.Errorf("Get %d, want 1000", s.Get())
	}
}
