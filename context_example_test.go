package wtype_test

import (
	"fmt"

	"github.com/wuchieh/wtype"
)

func ExampleNewContext() {
	ctx := wtype.NewContext(0)

	type tempCtx = wtype.Context[int]

	ctx = wtype.AddHandler(ctx,
		func(c *tempCtx) {
			fmt.Println("1")
			c.Next()
			fmt.Println("3")
			fmt.Println("Data:", c.C)
		},
		func(c *tempCtx) {
			fmt.Println("2")
			c.C++
		},
	)

	ctx.Do()
	fmt.Println("==========")
	ctx.Do()

	// output:
	// 1
	// 2
	// 3
	// Data: 1
	// ==========
	// 1
	// 2
	// 3
	// Data: 1
}
