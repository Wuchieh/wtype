package wtype_test

import (
	"encoding/json"
	"testing"

	"github.com/wuchieh/wtype"
)

func TestSyncMap(t *testing.T) {
	m := wtype.NewSyncMap[string, int]()
	m.Store("a", 1)
	m.Store("b", 2)
	m.Range(func(key string, value int) bool {
		m.Store(key, value+1)
		return true
	})
	m.Range(func(key string, value int) bool {
		t.Log(key, value)
		return true
	})
	actual, loaded := m.LoadOrStore("a", 1)
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal("json.Marshal Error:", err)
	} else {
		t.Log(string(b))
	}
	t.Log(actual, loaded)
}
