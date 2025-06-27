package wtype

// Set is a generic, non-thread-safe set implementation.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet creates a new empty Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

// Add adds an element to the set.
func (s *Set[T]) Add(data T) {
	s.m[data] = struct{}{}
}

// Get returns all elements in the set as a slice.
//
//	The order of elements is not guaranteed.
func (s *Set[T]) Get() []T {
	ret := make([]T, 0, len(s.m))
	for k := range s.m {
		ret = append(ret, k)
	}
	return ret
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Remove removes an element from the set.
func (s *Set[T]) Remove(data T) {
	delete(s.m, data)
}

// Contains checks if an element exists in the set.
func (s *Set[T]) Contains(data T) bool {
	_, ok := s.m[data]
	return ok
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

// Range iterates over the set and calls f for each element.
//
//	If f returns false, the iteration stops.
func (s *Set[T]) Range(f func(T) bool) {
	for k := range s.m {
		if !f(k) {
			break
		}
	}
}
