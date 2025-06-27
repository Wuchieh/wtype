package wtype_test

import (
	"fmt"
	"sort"
	"strings"

	"github.com/wuchieh/wtype"
)

// ExampleNewSet demonstrates creating a new set
func ExampleNewSet() {
	// Create a new string set
	stringSet := wtype.NewSet[string]()
	fmt.Printf("New string set length: %d\n", stringSet.Len())

	// Create a new int set
	intSet := wtype.NewSet[int]()
	fmt.Printf("New int set length: %d\n", intSet.Len())

	// Output:
	// New string set length: 0
	// New int set length: 0
}

// ExampleSet_Add demonstrates adding elements to a set
func ExampleSet_Add() {
	set := wtype.NewSet[string]()

	// Add elements
	set.Add("apple")
	set.Add("banana")
	set.Add("cherry")

	fmt.Printf("Set length after adding 3 elements: %d\n", set.Len())

	// Adding duplicate element (no effect)
	set.Add("apple")
	fmt.Printf("Set length after adding duplicate: %d\n", set.Len())

	// Output:
	// Set length after adding 3 elements: 3
	// Set length after adding duplicate: 3
}

// ExampleSet_Contains demonstrates checking if elements exist in the set
func ExampleSet_Contains() {
	set := wtype.NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	fmt.Printf("Contains 2: %t\n", set.Contains(2))
	fmt.Printf("Contains 5: %t\n", set.Contains(5))
	fmt.Printf("Contains 1: %t\n", set.Contains(1))

	// Output:
	// Contains 2: true
	// Contains 5: false
	// Contains 1: true
}

// ExampleSet_Remove demonstrates removing elements from a set
func ExampleSet_Remove() {
	set := wtype.NewSet[string]()
	set.Add("red")
	set.Add("green")
	set.Add("blue")

	fmt.Printf("Before removal - Length: %d, Contains 'green': %t\n",
		set.Len(), set.Contains("green"))

	// Remove an existing element
	set.Remove("green")
	fmt.Printf("After removing 'green' - Length: %d, Contains 'green': %t\n",
		set.Len(), set.Contains("green"))

	// Remove a non-existing element (no effect)
	set.Remove("yellow")
	fmt.Printf("After removing non-existing 'yellow' - Length: %d\n", set.Len())

	// Output:
	// Before removal - Length: 3, Contains 'green': true
	// After removing 'green' - Length: 2, Contains 'green': false
	// After removing non-existing 'yellow' - Length: 2
}

// ExampleSet_Get demonstrates getting all elements from a set
func ExampleSet_Get() {
	set := wtype.NewSet[int]()
	set.Add(3)
	set.Add(1)
	set.Add(4)
	set.Add(2)

	elements := set.Get()
	// Sort for consistent output since set order is not guaranteed
	sort.Ints(elements)
	fmt.Printf("Set elements: %v\n", elements)
	fmt.Printf("Number of elements: %d\n", len(elements))

	// Output:
	// Set elements: [1 2 3 4]
	// Number of elements: 4
}

// ExampleSet_Len demonstrates getting the length of a set
func ExampleSet_Len() {
	set := wtype.NewSet[string]()

	fmt.Printf("Empty set length: %d\n", set.Len())

	set.Add("first")
	fmt.Printf("After adding 1 element: %d\n", set.Len())

	set.Add("second")
	set.Add("third")
	fmt.Printf("After adding 3 elements total: %d\n", set.Len())

	// Output:
	// Empty set length: 0
	// After adding 1 element: 1
	// After adding 3 elements total: 3
}

// ExampleSet_Clear demonstrates clearing all elements from a set
func ExampleSet_Clear() {
	set := wtype.NewSet[int]()
	set.Add(10)
	set.Add(20)
	set.Add(30)

	fmt.Printf("Before clear - Length: %d\n", set.Len())

	set.Clear()
	fmt.Printf("After clear - Length: %d\n", set.Len())
	fmt.Printf("Contains 10 after clear: %t\n", set.Contains(10))

	// Output:
	// Before clear - Length: 3
	// After clear - Length: 0
	// Contains 10 after clear: false
}

// ExampleSet_Range demonstrates iterating over set elements
func ExampleSet_Range() {
	set := wtype.NewSet[string]()
	set.Add("apple")
	set.Add("banana")
	set.Add("cherry")

	fmt.Println("All elements:")
	set.Range(func(element string) bool {
		fmt.Printf("- %s\n", element)
		return true // Continue iteration
	})

	fmt.Println("\nFirst 2 elements:")
	count := 0
	set.Range(func(element string) bool {
		fmt.Printf("- %s\n", element)
		count++
		return count < 2 // Stop after 2 elements
	})

	// Output will vary due to map iteration order, but structure will be:
	// All elements:
	// - apple
	// - banana
	// - cherry
	//
	// First 2 elements:
	// - apple
	// - banana
}

// ExampleSet_Range_conditionalStop demonstrates stopping iteration early
func ExampleSet_Range_conditionalStop() {
	set := wtype.NewSet[int]()
	for i := 1; i <= 10; i++ {
		set.Add(i)
	}

	fmt.Println("Numbers until we find one > 5:")
	set.Range(func(num int) bool {
		fmt.Printf("%d ", num)
		return num <= 5 // Stop when we find a number > 5
	})
	fmt.Println()

	// Output will vary due to map iteration order, but will stop early
}

// ExampleSet_withDifferentTypes demonstrates sets with different data types
func ExampleSet_withDifferentTypes() {
	// String set
	stringSet := wtype.NewSet[string]()
	stringSet.Add("Go")
	stringSet.Add("Python")
	stringSet.Add("Java")
	fmt.Printf("Programming languages: %d\n", stringSet.Len())

	// Int set
	intSet := wtype.NewSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)
	fmt.Printf("Numbers: %d\n", intSet.Len())

	// Float set
	floatSet := wtype.NewSet[float64]()
	floatSet.Add(3.14)
	floatSet.Add(2.71)
	floatSet.Add(1.41)
	fmt.Printf("Constants: %d\n", floatSet.Len())

	// Output:
	// Programming languages: 3
	// Numbers: 3
	// Constants: 3
}

// ExampleSet_setOperations demonstrates common set operations
func ExampleSet_setOperations() {
	// Create two sets
	set1 := wtype.NewSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2 := wtype.NewSet[int]()
	set2.Add(3)
	set2.Add(4)
	set2.Add(5)

	// Union: elements in either set
	union := wtype.NewSet[int]()
	set1.Range(func(element int) bool {
		union.Add(element)
		return true
	})
	set2.Range(func(element int) bool {
		union.Add(element)
		return true
	})

	// Intersection: elements in both sets
	intersection := wtype.NewSet[int]()
	set1.Range(func(element int) bool {
		if set2.Contains(element) {
			intersection.Add(element)
		}
		return true
	})

	unionElements := union.Get()
	sort.Ints(unionElements)
	fmt.Printf("Union: %v\n", unionElements)

	intersectionElements := intersection.Get()
	sort.Ints(intersectionElements)
	fmt.Printf("Intersection: %v\n", intersectionElements)

	// Output:
	// Union: [1 2 3 4 5]
	// Intersection: [3]
}

// ExampleSet_uniqueWords demonstrates practical usage for finding unique words
func ExampleSet_uniqueWords() {
	text := "the quick brown fox jumps over the lazy dog the fox is quick"
	words := strings.Fields(text)

	// Use set to find unique words
	uniqueWords := wtype.NewSet[string]()
	for _, word := range words {
		uniqueWords.Add(word)
	}

	fmt.Printf("Total words: %d\n", len(words))
	fmt.Printf("Unique words: %d\n", uniqueWords.Len())

	// Get unique words and sort for consistent output
	unique := uniqueWords.Get()
	sort.Strings(unique)
	fmt.Printf("Unique word list: %v\n", unique)

	// Output:
	// Total words: 13
	// Unique words: 9
	// Unique word list: [brown dog fox is jumps lazy over quick the]
}

// ExampleSet_emptySet demonstrates working with empty sets
func ExampleSet_emptySet() {
	set := wtype.NewSet[string]()

	fmt.Printf("Empty set length: %d\n", set.Len())
	fmt.Printf("Contains anything: %t\n", set.Contains("anything"))

	// Get from empty set
	elements := set.Get()
	fmt.Printf("Elements in empty set: %v\n", elements)

	// Range over empty set
	fmt.Println("Ranging over empty set:")
	set.Range(func(element string) bool {
		fmt.Printf("This won't print: %s\n", element)
		return true
	})
	fmt.Println("Range completed")

	// Output:
	// Empty set length: 0
	// Contains anything: false
	// Elements in empty set: []
	// Ranging over empty set:
	// Range completed
}
