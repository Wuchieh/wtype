package wtype

import (
	"reflect"
	"testing"
)

func TestGorm(t *testing.T) {
	var uintSlice GormSlice[uint]

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
