package wtype_test

import (
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

func TestCache(t *testing.T) {
	c := wtype.NewCache(0, "123")

	t.Log(c.Get())
	time.Sleep(2 * time.Second)
	t.Log(c.Get())
}
