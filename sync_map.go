package wtype

import (
	"encoding/json"
	"sync"
)

type SyncMap[K comparable, V any] struct {
	m sync.Map
}

func (s *SyncMap[K, V]) MarshalJSON() ([]byte, error) {
	m := make(map[K]V)
	s.m.Range(func(key, value any) bool {
		m[key.(K)] = value.(V)
		return true
	})
	return json.Marshal(m)
}

func (s *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := s.m.Load(key)
	if ok {
		value = v.(V)
	}
	return value, ok
}

func (s *SyncMap[K, V]) Store(key K, value V) {
	s.m.Store(key, value)
}

func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := s.m.LoadOrStore(key, value)
	actual = v.(V)
	return actual, loaded
}

func (s *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := s.m.LoadAndDelete(key)
	value = v.(V)
	return value, loaded
}

func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, loaded := s.m.Swap(key, value)
	previous = v.(V)
	return previous, loaded
}

func (s *SyncMap[K, V]) CompareAndSwap(key K, old V, new V) (swapped bool) {
	swapped = s.m.CompareAndSwap(key, old, new)
	return swapped
}

func (s *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	deleted = s.m.CompareAndDelete(key, old)
	return deleted
}

func (s *SyncMap[K, V]) Range(f func(key K, value V) (shouldContinue bool)) {
	s.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (s *SyncMap[K, V]) Clear() {
	s.m.Clear()
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{}
}
