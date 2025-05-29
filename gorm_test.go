package wtype

import "testing"

func TestGorm(t *testing.T) {
	var uintSlice GormSlice[uint]

	t.Log(uintSlice.Value())

	for i := 0; i < 20; i++ {
		uintSlice = append(uintSlice, uint(i))
	}

	t.Log(uintSlice.Value())

	uintSlice.Scan([]byte("[10,20,30]"))

	t.Log(uintSlice)
}
