package wtype_test

import (
	"fmt"

	"github.com/wuchieh/wtype"
)

func ExampleNewSafeSet() {
	ids := wtype.NewSafeSet(1, 2, 3, 1, 2, 3)
	ids.Range(func(i int) bool {
		ids.Add(i + 2)
		return true
	})
	ids.Remove(4)

	fmt.Println(ids.Len())
	fmt.Println(ids.Contains(4))
	fmt.Println(ids.Contains(5))

	// output:
	// 4
	// false
	// true
}
