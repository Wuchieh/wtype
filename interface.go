package wtype

import "time"

type ICache[T any] interface {
	Set(T)
	Get() T
	SetDuration(time.Duration)
	ResetTimer()
}
