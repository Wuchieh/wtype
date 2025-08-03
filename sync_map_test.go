package wtype_test

import (
	"github.com/wuchieh/wtype"
	"testing"
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
	t.Log(actual, loaded)
}
