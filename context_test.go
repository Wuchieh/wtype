package wtype_test

import (
	"testing"
	"time"
	"unsafe"

	"github.com/wuchieh/wtype"
)

func TestContext(t *testing.T) {
	c := wtype.NewContext(1)
	var p2 unsafe.Pointer
	var ok bool
	p := unsafe.Pointer(&c)

	c = wtype.AddHandler(c,
		func(c *wtype.Context[int]) {
			done := c.Done()
			go func() {
				<-done
				ok = true
			}()
		},
		func(c *wtype.Context[int]) {
			p2 = unsafe.Pointer(c)
		},
		func(c *wtype.Context[int]) {
			c.C++
		},
		func(c *wtype.Context[int]) {
			c.C++
		},
		func(c *wtype.Context[int]) {

			if p == p2 || c.C != 3 {
				t.Error("context data error")
			}
		},
	)
	c.Do()
	if c.C == 3 {
		t.Error("src context data error")
	}
	time.Sleep(time.Millisecond)
	if !ok {
		t.Error("context error no ok")
	}
}
