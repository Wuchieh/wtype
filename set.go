package wtype

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(data T) {
	s.m[data] = struct{}{}
}

func (s *Set[T]) Get() []T {
	ret := make([]T, len(s.m))
	i := 0
	for k := range s.m {
		ret[i] = k
		i++
	}
	return ret
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Remove(data T) {
	delete(s.m, data)
}
