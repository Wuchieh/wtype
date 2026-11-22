package wtype

import "time"

// ICache defines the behavior of cache implementations.
type ICache[T any] interface {
	// Set stores the given value.
	Set(T)
	// Get returns the stored value.
	Get() T
	// SetDuration sets the cache duration.
	SetDuration(time.Duration)
	// ResetTimer resets the cache timer.
	ResetTimer()
	// StopTimer stops the cache timer.
	StopTimer()
}

// ISet defines the behavior of set implementations.
type ISet[T comparable] interface {
	// Add adds an element to the set.
	Add(T)
	// Get returns all elements as a slice.
	Get() []T
	// Len returns the number of elements in the set.
	Len() int
	// Remove removes an element from the set.
	Remove(T)
	// Clear removes all elements from the set.
	Clear()
	// Range iterates over the set and stops when f returns false.
	Range(func(T) bool)
	// Contains reports whether an element exists in the set.
	Contains(T) bool
	// Values returns all elements as a slice.
	Values() []T
}

// IContext defines the behavior of handler-chain contexts.
type IContext interface {
	// Next executes the remaining handlers.
	Next()
	// Abort prevents the remaining handlers from being executed.
	Abort()
	// IsAborted reports whether the context was aborted.
	IsAborted() bool
	// Get returns the value stored for the given key.
	Get(string) (any, bool)
	// Set stores a value for the given key.
	Set(string, any)
}

// IContextDoBefore defines a context that prepares a copy before execution.
type IContextDoBefore interface {
	// DoBefore returns the context to execute.
	DoBefore() IContext
}

// IMap defines the behavior of concurrent map implementations.
type IMap[K comparable, V any] interface {
	// Load returns the value stored for key.
	Load(key K) (value V, ok bool)
	// Store sets the value for key.
	Store(key K, value V)
	// LoadOrStore returns the existing value if present, otherwise it stores and
	// returns value.
	LoadOrStore(key K, value V) (actual V, loaded bool)
	// LoadAndDelete returns the value stored for key and deletes it.
	LoadAndDelete(key K) (value V, loaded bool)
	// Delete removes key from the map.
	Delete(K)
	// Swap stores value for key and returns the previous value.
	Swap(key K, value V) (previous V, loaded bool)
	// CompareAndSwap swaps old with new when the current value equals old.
	CompareAndSwap(key K, old, new V) (swapped bool)
	// CompareAndDelete deletes key when its current value equals old.
	CompareAndDelete(key K, old V) (deleted bool)
	// Range calls f for each entry and stops when f returns false.
	Range(func(key K, value V) (shouldContinue bool))
	// Clear removes all entries from the map.
	Clear()
}
