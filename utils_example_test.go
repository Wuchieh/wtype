package wtype_test

import (
	"fmt"

	"github.com/wuchieh/wtype"
)

func ExampleFallback() {
	r1 := wtype.Fallback(0, 5, 10)                         // returns 5
	r2 := wtype.Fallback("", "hello", "hi")                // returns "hello"
	r3 := wtype.Fallback(false, false, true)               // returns true
	r4 := wtype.Fallback(0, 0, 0)                          // returns 0
	r5 := wtype.Fallback([]int(nil), []int{}, []int{1, 2}) // returns []int{} (non-nil, not zero)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)

	//output:
	// 5
	// hello
	// true
	// 0
	// []
}
