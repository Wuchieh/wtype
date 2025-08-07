package wtype

import (
	"reflect"
	"strings"
)

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

	trimValue(rv)
}

func trimValue(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Struct:
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			fieldType := rt.Field(i)

			if field.CanSet() {
				trimValue(field)
			} else if fieldType.Anonymous {
				trimValue(field)
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
			trimValue(rv.Elem())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			trimValue(rv.Index(i))
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
