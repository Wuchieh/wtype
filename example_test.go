package wtype_test

import (
	"fmt"
	"time"

	"github.com/wuchieh/wtype"
)

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
