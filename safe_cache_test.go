package wtype_test

import (
	"sync"
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

func TestSafeCache(t *testing.T) {
	s := wtype.NewSafeCache(time.Second, 0)

	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(index int) {
			defer wg.Done()
			s.Use(func(data int) int {
				t.Logf("data: %d, index: %d", data, index)
				return data + 1
			})
		}(i)
	}
	wg.Wait()
}
