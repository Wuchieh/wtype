package wtype_test

import (
	"fmt"
	"time"

	"github.com/wuchieh/wtype"
)

// ExampleNewCache demonstrates creating a new cache with different configurations
func ExampleNewCache() {
	// Create a cache that expires after 5 seconds
	cache := wtype.NewCache[string](5 * time.Second)
	fmt.Printf("Empty cache: %q\n", cache.Get())

	// Create a cache with initial data
	cacheWithData := wtype.NewCache[int](3*time.Second, 42)
	fmt.Printf("Cache with initial data: %d\n", cacheWithData.Get())

	// Create a cache that never expires
	permanentCache := wtype.NewCache[string](0)
	permanentCache.Set("permanent data")
	fmt.Printf("Permanent cache: %q\n", permanentCache.Get())

	// Output:
	// Empty cache: ""
	// Cache with initial data: 42
	// Permanent cache: "permanent data"
}

// ExampleNewCache_neverExpire demonstrates creating a cache that never expires
func ExampleNewCache_neverExpire() {
	// Create cache with zero duration (never expires)
	cache := wtype.NewCache[string](0, "permanent data")
	fmt.Println(cache.Get())

	time.Sleep(100 * time.Millisecond)
	fmt.Println(cache.Get())

	// Output:
	// permanent data
	// permanent data
}

// ExampleCache_Get demonstrates getting data from the cache
func ExampleCache_Get() {
	cache := wtype.NewCache[int](1 * time.Second)

	// Get from empty cache (zero value)
	fmt.Printf("Empty cache: %d\n", cache.Get())

	// Set and get data
	cache.Set(100)
	fmt.Printf("With data: %d\n", cache.Get())

	// Output:
	// Empty cache: 0
	// With data: 100
}

// ExampleCache_Set demonstrates setting data in the cache
func ExampleCache_Set() {
	cache := wtype.NewCache[string](2 * time.Second)

	cache.Set("Hello, World!")
	fmt.Printf("After set: %q\n", cache.Get())

	// Setting new data resets the timer
	cache.Set("Updated data")
	fmt.Printf("After update: %q\n", cache.Get())

	// Output:
	// After set: "Hello, World!"
	// After update: "Updated data"
}

// ExampleCache_SetDuration demonstrates changing cache duration
func ExampleCache_SetDuration() {
	cache := wtype.NewCache[string](1 * time.Second)
	cache.Set("test data")

	// Change duration to 5 seconds
	cache.SetDuration(5 * time.Second)
	fmt.Printf("Duration changed, data: %q\n", cache.Get())

	// Change to permanent (0 duration)
	cache.SetDuration(0)
	fmt.Printf("Permanent cache, data: %q\n", cache.Get())

	// Output:
	// Duration changed, data: "test data"
	// Permanent cache, data: "test data"
}

// ExampleCache_StopTimer demonstrates stopping the cache timer
func ExampleCache_StopTimer() {
	cache := wtype.NewCache[string](100 * time.Millisecond)
	cache.Set("important data")

	// Stop the timer to prevent expiration
	cache.StopTimer()
	fmt.Printf("Timer stopped, data: %q\n", cache.Get())

	// Wait beyond the original expiration time
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("After wait, data still: %q\n", cache.Get())

	// Output:
	// Timer stopped, data: "important data"
	// After wait, data still: "important data"
}

// ExampleCache_ResetTimer demonstrates resetting the cache timer
func ExampleCache_ResetTimer() {
	cache := wtype.NewCache[string](100 * time.Millisecond)
	cache.Set("data")

	// Wait a bit, then reset timer
	time.Sleep(50 * time.Millisecond)
	cache.ResetTimer()
	fmt.Printf("Timer reset, data: %q\n", cache.Get())

	// Data should still be available after original time would have expired
	time.Sleep(75 * time.Millisecond)
	fmt.Printf("After partial wait, data: %q\n", cache.Get())

	// Output:
	// Timer reset, data: "data"
	// After partial wait, data: "data"
}
