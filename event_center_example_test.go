package wtype_test

import (
	"fmt"

	"github.com/wuchieh/wtype"
)

func ExampleNewEventCenter() {
	center := wtype.NewEventCenter()

	creatFunc := func(i int) func(data ...any) {
		return func(data ...any) {
			fmt.Println(i)
		}
	}

	run1 := creatFunc(1)
	run2 := creatFunc(2)
	run3 := creatFunc(3)
	run4 := creatFunc(4)

	center.On("run", run1)
	center.On("run", run2)
	center.On("run", run3)
	center.On("run", run4)

	center.Emit("run")
	center.Off("run", run3)
	fmt.Println("-----")
	center.Emit("run")

	// Output:
	// 1
	// 2
	// 3
	// 4
	// -----
	// 1
	// 2
	// 4
}
