package wtype

import "sync"

// SafeSet is a thread-safe version of Set.
type SafeSet[T comparable] struct {
	mx sync.RWMutex
	s  Set[T]
}

// MarshalJSON implementation json.Marshal
// nolint:govet
func (s SafeSet[T]) MarshalJSON() ([]byte, error) {
	return s.s.MarshalJSON()
}

// UnmarshalJSON implementation json.Unmarshal
func (s *SafeSet[T]) UnmarshalJSON(bytes []byte) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.s.UnmarshalJSON(bytes)
}

// Values returns all elements in the set as a slice.
//
//	The order of elements is not guaranteed.
func (s *SafeSet[T]) Values() []T {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.s.Values()
}

// Add adds an element to the set.
func (s *SafeSet[T]) Add(data T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.s.Add(data)
}

// Get returns all elements in the set as a slice.
// The order of elements is not guaranteed.
func (s *SafeSet[T]) Get() []T {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.s.Get()
}

// Len returns the number of elements in the set.
func (s *SafeSet[T]) Len() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.s.Len()
}

// Remove removes an element from the set.
func (s *SafeSet[T]) Remove(data T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.s.Remove(data)
}

// Contains checks if an element exists in the set.
func (s *SafeSet[T]) Contains(data T) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.s.Contains(data)
}

// Clear removes all elements from the set.
func (s *SafeSet[T]) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.s.Clear()
}

// Range iterates over the set and calls f for each element.
// If f returns false, the iteration stops.
// The iteration is performed under a read lock.
func (s *SafeSet[T]) Range(f func(T) bool) {
	s.mx.RLock()
	cp := s.s.Values()
	s.mx.RUnlock()
	for _, v := range cp {
		if !f(v) {
			break
		}
	}
}

// SortValues sort the set values
func (s *SafeSet[T]) SortValues(cmp func(a T, b T) bool) []T {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.s.SortValues(cmp)
}

// NewSafeSet creates a new empty SafeSet.
func NewSafeSet[T comparable](val ...T) *SafeSet[T] {
	s := SafeSet[T]{
		s: *NewSet[T](val...),
	}
	return &s
}
