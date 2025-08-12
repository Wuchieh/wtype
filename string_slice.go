package wtype

import "strings"

type SliceString []String

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

func (s *SliceString) ToString() []string {
	return SliceConvert(*s, func(t String) string {
		return t.ToString()
	})
}

func NewSliceString[T string | String](s ...T) SliceString {
	return SliceConvert(s, func(t T) String {
		return *NewString(string(t))
	})
}
