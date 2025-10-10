package wtype

import "strings"

type String string

func (s *String) String() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

func (s *String) ToString() string {
	return s.String()
}

func (s *String) Trim() *String {
	t := s.String()
	return NewString(strings.TrimSpace(t))
}

func (s *String) Repeat(count int) *String {
	t := s.String()
	return NewString(strings.Repeat(t, count))
}

func (s *String) Replace(old, new string, n int) *String {
	t := s.String()
	return NewString(strings.Replace(t, old, new, n))
}

func (s *String) ReplaceAll(old, new string) *String {
	t := s.String()
	return NewString(strings.ReplaceAll(t, old, new))
}

func (s *String) Slice(start int, end ...int) *String {
	return NewString(StringSlice(s.String(), start, end...))
}

func (s *String) Split(sep ...string) SliceString {
	se := ""

	if len(sep) > 0 {
		se = sep[0]
	}
	return NewSliceString(strings.Split(s.String(), se)...)
}

func (s *String) Contains(substr string) bool {
	return strings.Contains(s.String(), substr)
}

func (s *String) Includes(substr string) bool {
	return s.Contains(substr)
}

func (s *String) ToLower() *String {
	return NewString(strings.ToLower(s.String()))
}

func (s *String) ToUpper() *String {
	return NewString(strings.ToUpper(s.String()))
}

func (s *String) Len() int {
	return StringLen(s.ToString())
}

func (s *String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.String(), suffix)
}

func (s *String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.String(), prefix)
}

func NewString[T ~string](s T) *String {
	S := String(s)
	return &S
}
