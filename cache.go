package wtype

import "time"

type Cache[T any] struct {
	data T
	d    time.Duration
	t    *time.Timer
}

// setTimer sets the timer of the cache.
func (c *Cache[T]) setTimer() {
	if c.t != nil {
		c.t.Stop()
	}

	d := c.d
	if d == 0 {
		return
	}

	c.t = time.AfterFunc(d, c.resetData)
}

// resetData resets the data of the cache.
func (c *Cache[T]) resetData() {
	c.data = *new(T)
}

// SetDuration sets the duration of the cache.
func (c *Cache[T]) SetDuration(d time.Duration) {
	c.d = d
}

// Set sets the data of the cache.
func (c *Cache[T]) Set(data T) {
	c.data = data
	c.setTimer()
}

func (c *Cache[T]) Get() T {
	return c.data
}

// NewCache creates a new cache.
// If d is 0, the cache will never expire.
func NewCache[T any](d time.Duration, data ...T) *Cache[T] {
	c := &Cache[T]{}
	c.resetData()
	c.SetDuration(d)
	if len(data) > 0 {
		c.Set(data[0])
	}
	return c
}
