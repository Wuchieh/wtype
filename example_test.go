package wtype_test

import (
	"fmt"
	"github.com/wuchieh/wtype"
	"time"
)

func ExampleCache_ResetTimer() {
	c := wtype.NewCache[int](time.Second)
	c.Set(1)
	time.Sleep(time.Millisecond)
	c.ResetTimer()
	time.Sleep(time.Millisecond)
	fmt.Println(c.Get())
	time.Sleep(time.Second)
	fmt.Println(c.Get())

	// output:
	// 1
	// 0
}
