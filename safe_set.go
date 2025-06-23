package wtype

import "sync"

type SafeSet[T comparable] struct {
	mx sync.RWMutex
	s  Set[T]
}

func (s *SafeSet[T]) Add(data T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.s.Add(data)
}

func (s *SafeSet[T]) Get() []T {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.s.Get()
}

func (s *SafeSet[T]) Len() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.s.Len()
}

func (s *SafeSet[T]) Remove(data T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.s.Remove(data)
}

func NewSafeSet[T comparable]() *SafeSet[T] {
	return &SafeSet[T]{
		s: Set[T]{m: make(map[T]struct{})},
	}
}
