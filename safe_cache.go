package wtype

import (
	"sync"
	"time"
)

type SafeCache[T any] struct {
	cache Cache[T]
	mutex sync.RWMutex
}

// StopTimer stops the timer of the cache.
//
//	The data will be retained permanently.
func (s *SafeCache[T]) StopTimer() {
	s.cache.StopTimer()
}

// ResetTimer resets the timer of the cache.
func (s *SafeCache[T]) ResetTimer() {
	s.cache.ResetTimer()
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

// Use uses the data of the cache.
//
//	The data will be updated after the function is called.
func (s *SafeCache[T]) Use(f func(T) T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.cache.Set(f(s.cache.Get()))
}

// Use2 uses the data of the cache.
//
//	If the data is not set, the function will be called and the result will be set.
func (s *SafeCache[T]) Use2(f func(T) (T, error)) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	nd, err := f(s.cache.Get())
	if err != nil {
		return err
	}
	s.cache.Set(nd)
	return nil
}

// NewSafeCache creates a new safe cache.
func NewSafeCache[T any](d time.Duration, data ...T) *SafeCache[T] {
	c := NewCache(d, data...)
	s := &SafeCache[T]{
		cache: *c,
	}
	return s
}
