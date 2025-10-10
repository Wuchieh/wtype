package wtype_test

import (
	"encoding/json"
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

func TestSafeSet_JSON(t *testing.T) {
	s := wtype.NewSafeSet(1, 2, 3, 4, 5)
	data, err := json.Marshal(s)
	if err != nil {
		t.Fatal("SafeSet json.Marshal Fail:", err)
	}

	var s2 wtype.SafeSet[uint64]
	err = json.Unmarshal(data, &s2)
	if err != nil {
		t.Fatal("SafeSet json.Unmarshal Fail:", err)
	}

	if s.Len() != s2.Len() {
		t.Fatal("SafeSet Len Fail")
	}

	s2.Range(func(value uint64) bool {
		if s.Contains(int(value)) {
			return true
		}
		t.Error("SafeSet Range Fail:", value)
		return false
	})
}
