package wtype

import "strings"

// String is a string type with chainable helper methods.
type String string

// String returns the string value, or an empty string for a nil receiver.
func (s *String) String() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

// ToString returns the string value. It is an alias for String.
func (s *String) ToString() string {
	return s.String()
}

// Trim returns a copy with leading and trailing whitespace removed.
func (s *String) Trim() *String {
	t := s.String()
	return NewString(strings.TrimSpace(t))
}

// Repeat returns a copy repeated count times.
func (s *String) Repeat(count int) *String {
	t := s.String()
	return NewString(strings.Repeat(t, count))
}

// Replace returns a copy with the first n occurrences of old replaced by new.
func (s *String) Replace(old, new string, n int) *String {
	t := s.String()
	return NewString(strings.Replace(t, old, new, n))
}

// ReplaceAll returns a copy with all occurrences of old replaced by new.
func (s *String) ReplaceAll(old, new string) *String {
	t := s.String()
	return NewString(strings.ReplaceAll(t, old, new))
}

// Slice returns the substring between start and the optional end index, using
// rune offsets.
func (s *String) Slice(start int, end ...int) *String {
	return NewString(StringSlice(s.String(), start, end...))
}

// Split returns the substrings separated by sep, defaulting to an empty
// separator.
func (s *String) Split(sep ...string) SliceString {
	se := ""

	if len(sep) > 0 {
		se = sep[0]
	}
	return NewSliceString(strings.Split(s.String(), se)...)
}

// Contains reports whether substr is within the string.
func (s *String) Contains(substr string) bool {
	return strings.Contains(s.String(), substr)
}

// Includes reports whether substr is within the string. It is an alias for
// Contains.
func (s *String) Includes(substr string) bool {
	return s.Contains(substr)
}

// ToLower returns a copy converted to lower case.
func (s *String) ToLower() *String {
	return NewString(strings.ToLower(s.String()))
}

// ToUpper returns a copy converted to upper case.
func (s *String) ToUpper() *String {
	return NewString(strings.ToUpper(s.String()))
}

// Len returns the number of runes in the string.
func (s *String) Len() int {
	return StringLen(s.ToString())
}

// HasSuffix reports whether the string ends with suffix.
func (s *String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.String(), suffix)
}

// HasPrefix reports whether the string begins with prefix.
func (s *String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.String(), prefix)
}

// Count returns the number of non-overlapping occurrences of substr.
func (s *String) Count(substr string) int {
	return strings.Count(s.String(), substr)
}

// NewString creates a new String from s.
func NewString[T ~string](s T) *String {
	S := String(s)
	return &S
}
