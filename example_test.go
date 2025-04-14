package wtype

import (
	"fmt"
	"time"
)

func ExampleCache_ResetTimer() {
	c := NewCache[int](time.Second)
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
