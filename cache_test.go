package wtype

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(0, "123")

	t.Log(c.Get())
	time.Sleep(2 * time.Second)
	t.Log(c.Get())
}
