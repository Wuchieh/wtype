package wtype

import (
	"sync"
	"time"
)

type SafeCache[T any] struct {
	cache Cache[T]
	mutex sync.RWMutex
}

// SetDuration sets the duration of the cache.
func (s *SafeCache[T]) SetDuration(duration time.Duration) {
	s.cache.SetDuration(duration)
}

// Set sets the data of the cache.
func (s *SafeCache[T]) Set(data T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.cache.Set(data)
}

// Get gets the data of the cache.
func (s *SafeCache[T]) Get() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.cache.Get()
}

// NewSafeCache creates a new safe cache.
func NewSafeCache[T any](d time.Duration, data ...T) *SafeCache[T] {
	c := NewCache(d, data...)
	s := &SafeCache[T]{
		cache: *c,
	}
	return s
}
