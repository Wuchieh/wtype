package wtype

import "strings"

// SliceString is a slice of String values.
type SliceString []String

// Join concatenates the elements using sep, defaulting to an empty separator.
func (s *SliceString) Join(sep ...string) *String {
	if s == nil {
		return NewString("")
	}

	se := ""

	if len(sep) > 0 {
		se = sep[0]
	}

	return NewString(strings.Join(SliceConvert(*s, func(t String) string {
		return t.ToString()
	}), se))
}

// ToString converts the slice to a []string.
func (s *SliceString) ToString() []string {
	return SliceConvert(*s, func(t String) string {
		return t.ToString()
	})
}

// NewSliceString creates a SliceString from the given values.
func NewSliceString[T string | String](s ...T) SliceString {
	return SliceConvert(s, func(t T) String {
		return *NewString(string(t))
	})
}
