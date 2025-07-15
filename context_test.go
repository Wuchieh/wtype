package wtype_test

import (
	"fmt"
	"github.com/wuchieh/wtype"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := wtype.NewContext[int](0)
	ctx = wtype.AddHandler(ctx,
		func(c *wtype.Context[int]) {
			c.Next()
			fmt.Println(c.C)
		},
		func(c *wtype.Context[int]) {
			c.C++
		},
		func(c *wtype.Context[int]) {
			c.C++
		},
	)
	ctx.Next()
}
