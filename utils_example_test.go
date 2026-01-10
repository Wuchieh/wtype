package wtype_test

import (
	"context"
	"fmt"
	"sort"
	"time"

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

	// output:
	// 5
	// hello
	// true
	// 0
	// []
}

func ExampleSliceConvert2() {
	var data []int
	for i := 0; i < 10; i++ {
		data = append(data, i)
	}

	result := wtype.SliceConvert2(data, func(i int, v int) (int, bool) {
		if v%2 == 0 {
			return v, true
		}
		return 0, false
	})

	fmt.Println(result)

	// output:
	// [0 2 4 6 8]
}

func ExampleAssert() {
	a := any(int32(5))
	fmt.Println(wtype.Assert[int32](a))
	fmt.Println(wtype.Assert[int](a))

	// output:
	// 5 true
	// 0 false
}

func ExampleMapConvert() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := wtype.MapConvert(m, func(k string, v int) (int, string, bool) {
		if v%2 == 0 {
			return 0, "", false
		}
		return v, k, true
	})

	// map output order is random, so we sort keys for stable output
	keys := make([]int, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d:%s ", k, m2[k])
	}
	fmt.Println()

	// output:
	// 1:a 3:c
}

func ExampleMapReverse() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := wtype.MapReverse(m)

	// map output order is random, so we sort keys for stable output
	keys := make([]int, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d:%s ", k, m2[k])
	}
	fmt.Println()

	// output:
	// 1:a 2:b 3:c
}

func ExampleContextIsTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	fmt.Println(wtype.ContextIsTimeout(ctx))

	time.Sleep(200 * time.Millisecond)

	fmt.Println(wtype.ContextIsTimeout(ctx))

	// output:
	// false
	// true
}
