package wtype

import "time"

type ICache[T any] interface {
	Set(T)
	Get() T
	SetDuration(time.Duration)
	ResetTimer()
	StopTimer()
}

type ISet[T comparable] interface {
	Add(T)
	Get() []T
	Len() int
	Remove(T)
	Clear()
	Range(func(T) bool)
	Contains(T) bool
	Values() []T
}

type IContext interface {
	Next()
	Abort()
	IsAborted() bool
	Get(string) (any, bool)
	Set(string, any)
}

type IMap[K comparable, V any] interface {
	Load(key K) (value V, ok bool)
	Store(key K, value V)
	LoadOrStore(key K, value V) (actual V, loaded bool)
	LoadAndDelete(key K) (value V, loaded bool)
	Delete(K)
	Swap(key K, value V) (previous V, loaded bool)
	CompareAndSwap(key K, old, new V) (swapped bool)
	CompareAndDelete(key K, old V) (deleted bool)
	Range(func(key K, value V) (shouldContinue bool))
	Clear()
}
