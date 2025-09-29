package wtype_test

import (
	"reflect"
	"testing"

	"github.com/wuchieh/wtype"
)

func TestGorm(t *testing.T) {
	var uintSlice wtype.GormSlice[uint]

	t.Log(uintSlice.Value())

	for i := 0; i < 20; i++ {
		uintSlice = append(uintSlice, uint(i))
	}

	t.Log(uintSlice.Value())

	uintSlice.Scan([]byte("[10,20,30]"))

	t.Log(uintSlice)

	slices := uintSlice.ToSlice()

	t.Log(reflect.TypeOf(uintSlice), reflect.TypeOf(slices))
}
