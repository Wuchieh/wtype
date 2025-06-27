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

// ExampleCache_expiration demonstrates cache expiration behavior
func ExampleCache_expiration() {
	cache := wtype.NewCache[string](100 * time.Millisecond)
	cache.Set("will expire")

	fmt.Printf("Before expiration: %q\n", cache.Get())

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)
	fmt.Printf("After expiration: %q\n", cache.Get())

	// Output:
	// Before expiration: "will expire"
	// After expiration: ""
}

// ExampleCache_withStruct demonstrates using cache with custom struct
func ExampleCache_withStruct() {
	type User struct {
		Name string
		Age  int
	}

	cache := wtype.NewCache[User](1 * time.Second)

	// Set struct data
	user := User{Name: "Alice", Age: 30}
	cache.Set(user)

	retrieved := cache.Get()
	fmt.Printf("User: %+v\n", retrieved)

	// Output:
	// User: {Name:Alice Age:30}
}

// ExampleCache_withPointer demonstrates using cache with pointer types
func ExampleCache_withPointer() {
	type Config struct {
		Setting string
		Value   int
	}

	cache := wtype.NewCache[*Config](1 * time.Second)

	// Initially nil
	fmt.Printf("Initial: %v\n", cache.Get())

	// Set pointer data
	config := &Config{Setting: "debug", Value: 1}
	cache.Set(config)

	retrieved := cache.Get()
	if retrieved != nil {
		fmt.Printf("Config: %+v\n", *retrieved)
	}

	// Output:
	// Initial: <nil>
	// Config: {Setting:debug Value:1}
}

// ExampleCache_multipleTypes demonstrates cache with different types
func ExampleCache_multipleTypes() {
	// String cache
	stringCache := wtype.NewCache[string](1*time.Second, "default")
	fmt.Printf("String cache: %q\n", stringCache.Get())

	// Int cache
	intCache := wtype.NewCache[int](1 * time.Second)
	intCache.Set(42)
	fmt.Printf("Int cache: %d\n", intCache.Get())

	// Slice cache
	sliceCache := wtype.NewCache[[]int](1 * time.Second)
	sliceCache.Set([]int{1, 2, 3})
	fmt.Printf("Slice cache: %v\n", sliceCache.Get())

	// Output:
	// String cache: "default"
	// Int cache: 42
	// Slice cache: [1 2 3]
}

// SafeCache Examples

// ExampleNewSafeCache demonstrates creating a new thread-safe cache
func ExampleNewSafeCache() {
	// Create a safe cache that expires after 5 seconds
	safeCache := wtype.NewSafeCache[string](5 * time.Second)
	fmt.Printf("Empty safe cache: %q\n", safeCache.Get())

	// Create a safe cache with initial data
	safeCacheWithData := wtype.NewSafeCache[int](3*time.Second, 42)
	fmt.Printf("Safe cache with initial data: %d\n", safeCacheWithData.Get())

	// Create a safe cache that never expires
	permanentSafeCache := wtype.NewSafeCache[string](0)
	permanentSafeCache.Set("permanent safe data")
	fmt.Printf("Permanent safe cache: %q\n", permanentSafeCache.Get())

	// Output:
	// Empty safe cache: ""
	// Safe cache with initial data: 42
	// Permanent safe cache: "permanent safe data"
}

// ExampleSafeCache_Set demonstrates setting data in the thread-safe cache
func ExampleSafeCache_Set() {
	safeCache := wtype.NewSafeCache[string](2 * time.Second)

	safeCache.Set("Hello, Safe World!")
	fmt.Printf("After set: %q\n", safeCache.Get())

	// Setting new data resets the timer
	safeCache.Set("Updated safe data")
	fmt.Printf("After update: %q\n", safeCache.Get())

	// Output:
	// After set: "Hello, Safe World!"
	// After update: "Updated safe data"
}

// ExampleSafeCache_Get demonstrates getting data from the thread-safe cache
func ExampleSafeCache_Get() {
	safeCache := wtype.NewSafeCache[int](1 * time.Second)

	// Get from empty cache (zero value)
	fmt.Printf("Empty safe cache: %d\n", safeCache.Get())

	// Set and get data
	safeCache.Set(100)
	fmt.Printf("With data: %d\n", safeCache.Get())

	// Output:
	// Empty safe cache: 0
	// With data: 100
}

// ExampleSafeCache_Use demonstrates using the Use method to update cache data
func ExampleSafeCache_Use() {
	safeCache := wtype.NewSafeCache[int](5 * time.Second)
	safeCache.Set(10)

	fmt.Printf("Before Use: %d\n", safeCache.Get())

	// Use function to double the value
	safeCache.Use(func(current int) int {
		return current * 2
	})

	fmt.Printf("After Use (doubled): %d\n", safeCache.Get())

	// Use function to add 5
	safeCache.Use(func(current int) int {
		return current + 5
	})

	fmt.Printf("After Use (added 5): %d\n", safeCache.Get())

	// Output:
	// Before Use: 10
	// After Use (doubled): 20
	// After Use (added 5): 25
}

// ExampleSafeCache_Use2 demonstrates using the Use2 method with error handling
func ExampleSafeCache_Use2() {
	safeCache := wtype.NewSafeCache[string](5 * time.Second)
	safeCache.Set("hello")

	fmt.Printf("Before Use2: %q\n", safeCache.Get())

	// Use2 function to convert to uppercase
	err := safeCache.Use2(func(current string) (string, error) {
		if current == "" {
			return "", fmt.Errorf("empty string")
		}
		return fmt.Sprintf("%s WORLD", current), nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After Use2 (success): %q\n", safeCache.Get())
	}

	// Use2 function that returns an error
	safeCache.Set("")
	err = safeCache.Use2(func(current string) (string, error) {
		if current == "" {
			return "", fmt.Errorf("cannot process empty string")
		}
		return current + " processed", nil
	})

	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		fmt.Printf("Data unchanged: %q\n", safeCache.Get())
	}

	// Output:
	// Before Use2: "hello"
	// After Use2 (success): "hello WORLD"
	// Error occurred: cannot process empty string
	// Data unchanged: ""
}

// ExampleSafeCache_concurrency demonstrates thread-safe operations
func ExampleSafeCache_concurrency() {
	safeCache := wtype.NewSafeCache[int](10 * time.Second)
	safeCache.Set(0)

	// Simulate concurrent access
	done := make(chan bool, 10)

	// Start 10 goroutines to increment the counter
	for i := 0; i < 10; i++ {
		go func() {
			safeCache.Use(func(current int) int {
				return current + 1
			})
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	fmt.Printf("Final count after concurrent operations: %d\n", safeCache.Get())

	// Output:
	// Final count after concurrent operations: 10
}

// ExampleSafeCache_timerOperations demonstrates timer-related operations
func ExampleSafeCache_timerOperations() {
	safeCache := wtype.NewSafeCache[string](100 * time.Millisecond)
	safeCache.Set("timer test")

	fmt.Printf("Initial data: %q\n", safeCache.Get())

	// Stop timer to prevent expiration
	safeCache.StopTimer()
	time.Sleep(150 * time.Millisecond)
	fmt.Printf("After stopping timer: %q\n", safeCache.Get())

	// Change duration and reset timer
	safeCache.SetDuration(200 * time.Millisecond)
	safeCache.ResetTimer()
	fmt.Printf("After resetting timer: %q\n", safeCache.Get())

	// Output:
	// Initial data: "timer test"
	// After stopping timer: "timer test"
	// After resetting timer: "timer test"
}

// ExampleSafeCache_withComplexType demonstrates SafeCache with complex data types
func ExampleSafeCache_withComplexType() {
	type UserStats struct {
		Name       string
		LoginCount int
		LastLogin  time.Time
	}

	safeCache := wtype.NewSafeCache[UserStats](5 * time.Second)

	// Initialize user stats
	initialStats := UserStats{
		Name:       "Alice",
		LoginCount: 1,
		LastLogin:  time.Date(2025, 6, 25, 0, 0, 0, 0, time.UTC),
	}
	safeCache.Set(initialStats)

	fmt.Printf("Initial stats: %+v\n", safeCache.Get())

	// Update login count using Use
	safeCache.Use(func(current UserStats) UserStats {
		current.LoginCount++
		current.LastLogin = time.Date(2025, 6, 27, 0, 0, 0, 0, time.UTC)
		return current
	})

	updated := safeCache.Get()
	fmt.Printf("Updated login count: %d\n", updated.LoginCount)
	fmt.Printf("Username: %s\n", updated.Name)

	// Output:
	// Initial stats: {Name:Alice LoginCount:1 LastLogin:2025-06-25 00:00:00 +0000 UTC}
	// Updated login count: 2
	// Username: Alice
}
