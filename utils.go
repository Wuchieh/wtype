package wtype

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/sync/singleflight"
)

var (
	shared = singleflight.Group{}
)

type SharedChanResult[T any] struct {
	singleflight.Result
	Val T
}

// DoShared executes the given function fn associated with the provided key.
// If multiple calls with the same key are made concurrently, fn will only be
// executed once, and the result will be shared among all callers.
func DoShared[T any](key string, fn func() (T, error)) (T, error) {
	var result T
	do, err, _ := shared.Do(key, func() (interface{}, error) {
		t, err := fn()
		if err != nil {
			return nil, err
		}
		return t, nil
	})

	if err != nil {
		return result, err
	}

	result = do.(T)
	return result, nil
}

// DoSharedChan is the channel-based variant of DoShared.
// It executes the given function fn associated with the provided key.
// If multiple calls with the same key are made concurrently, fn will only be
// executed once, and the result will be delivered through the returned channel.
func DoSharedChan[T any](key string, fn func() (T, error)) <-chan SharedChanResult[T] {
	doChan := shared.DoChan(key, func() (interface{}, error) {
		t, err := fn()
		if err != nil {
			return nil, err
		}
		return t, nil
	})

	result := make(chan SharedChanResult[T])

	go func() {
		data := <-doChan

		result <- SharedChanResult[T]{
			Result: data,
			Val:    data.Val.(T),
		}
	}()

	return result
}

// DoSharedForget removes the entry for the given key from the shared group,
// allowing subsequent calls with the same key to re-execute the function fn
// instead of receiving a shared result.
func DoSharedForget(key string) {
	shared.Forget(key)
}

// StructStringTrim removes leading and trailing whitespace from a struct
func StructStringTrim(v any) {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr {
		return
	}

	rv = rv.Elem()
	for rv.Kind() == reflect.Ptr &&
		!rv.IsNil() {
		rv = rv.Elem()
	}

	trimValue(rv, make(map[reflect.Value]struct{}))
}

func trimValue(rv reflect.Value, record map[reflect.Value]struct{}) {
	if _, ok := record[rv]; ok {
		return
	} else {
		record[rv] = struct{}{}
	}

	switch rv.Kind() {
	case reflect.Struct:
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			fieldType := rt.Field(i)

			if field.CanSet() {
				trimValue(field, record)
			} else if fieldType.Anonymous {
				trimValue(field, record)
			}
		}
	case reflect.Ptr:
		if !rv.IsNil() && rv.Elem().Kind() == reflect.String {
			if rv.Elem().CanSet() {
				rv.Elem().SetString(strings.TrimSpace(rv.Elem().String()))
			}
		} else if !rv.IsNil() &&
			rv.Elem().Kind() == reflect.Struct ||
			rv.Elem().Kind() == reflect.Slice {
			trimValue(rv.Elem(), record)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			trimValue(rv.Index(i), record)
		}
	case reflect.String:
		if rv.CanSet() {
			rv.SetString(strings.TrimSpace(rv.String()))
		}
	default:
		return
	}
}

// SliceToMap converts a slice to a map
func SliceToMap[T any, K comparable](slice []T, getKey func(int, T) K) map[K]T {
	m := make(map[K]T, len(slice))
	for i, v := range slice {
		m[getKey(i, v)] = v
	}
	return m
}

// StringSlice returns a substring of a string
func StringSlice(s string, start int, end ...int) string {
	runes := []rune(s)
	length := len(runes)

	if start < 0 {
		start = length + start
	}
	if start < 0 {
		start = 0
	}
	if start > length {
		start = length
	}

	e := length
	if len(end) > 0 {
		e = end[0]
		if e < 0 {
			e = length + e
		}
		if e < 0 {
			e = 0
		}
		if e > length {
			e = length
		}
	}

	if e <= start {
		return ""
	}

	return string(runes[start:e])
}

// SliceConvert converts a slice of type T to a slice of type K
func SliceConvert[T, K any](s []T, f func(T) K) []K {
	ret := make([]K, 0, len(s))
	for _, v := range s {
		ret = append(ret, f(v))
	}
	return ret
}

// SliceConvert2 converts a slice of type T to a slice of type K
//
//	If the function returns false, the element will be skipped.
func SliceConvert2[T, K any](s []T, f func(int, T) (K, bool)) []K {
	result := make([]K, 0, len(s))
	for i, v := range s {
		data, ok := f(i, v)
		if ok {
			result = append(result, data)
		}
	}
	return result
}

// SlicePointConvert converts a slice of type *T to a slice of type K
func SlicePointConvert[T any](s []T) []*T {
	return SliceConvert(s, func(v T) *T {
		return &v
	})
}

// SliceUnPointConvert converts a slice of type *T to a slice of type K
func SliceUnPointConvert[T any](s []*T) []T {
	return SliceConvert(s, func(v *T) T {
		if v == nil {
			return *new(T)
		}
		return *v
	})
}

func isZero(t any) bool {
	va := reflect.ValueOf(t)
	return va.IsZero()
}

// Fallback returns the first non-zero value from the given arguments.
// If all values are zero, it returns the zero value of type T.
func Fallback[T any](data ...T) T {
	var zero T
	for _, datum := range data {
		if !isZero(datum) {
			return datum
		}
	}
	return zero
}

// Stack returns a nicely formatted stack frame, skipping skip frames.
func Stack(skip int) []byte {
	// +1 是為了跳過當前 Stack 函數本身
	callers := make([]uintptr, 32)
	n := runtime.Callers(skip+2, callers) // +2 跳過 Callers 和 Stack 函數

	if n == 0 {
		return []byte("no stack available")
	}

	frames := runtime.CallersFrames(callers[:n])
	var buf strings.Builder

	for {
		frame, more := frames.Next()
		fmt.Fprintf(&buf, "%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)

		if !more {
			break
		}
	}

	return []byte(buf.String())
}
