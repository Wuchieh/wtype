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
}
