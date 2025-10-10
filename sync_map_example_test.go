package wtype_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wuchieh/wtype"
)

func ExampleNewSyncMap() {
	m := wtype.NewSyncMap[string, int]()
	m.Store("a", 1)
	m.Store("b", 2)
	m.Range(func(key string, value int) bool {
		m.Store(key, value+1)
		return true
	})
	m.Range(func(key string, value int) bool {
		fmt.Println(key, value)
		return true
	})

	actual, loaded := m.LoadOrStore("a", 1)
	fmt.Println(actual, loaded)

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal("json.Marshal Error:", err)
	}

	var m2 wtype.SyncMap[string, int]
	err = json.Unmarshal(b, &m2)
	if err != nil {
		log.Fatal("json.Unmarshal Error:", err)
	}

	a, _ := m2.Load("a")
	fmt.Println(a == 2)

	// output:
	// a 2
	// b 3
	// 2 true
	// true
}
