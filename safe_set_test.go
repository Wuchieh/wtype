package wtype

import (
	"sync"
	"testing"
)

func TestSafeSet(t *testing.T) {
	t.Run("BasicOperations", func(t *testing.T) {
		s := NewSafeSet[int]()
		if s.Len() != 0 {
			t.Error("New set should be empty")
		}

		s.Add(1)
		s.Add(2)
		s.Add(3)

		if s.Len() != 3 {
			t.Error("Set should have 3 elements")
		}

		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}

		s.Remove(2)
		if s.Contains(2) {
			t.Error("Set should not contain 2 after removal")
		}

		items := s.Get()
		if len(items) != 2 {
			t.Error("Get should return 2 items")
		}
	})

	t.Run("Concurrency", func(t *testing.T) {
		s := NewSafeSet[int]()
		var wg sync.WaitGroup
		const workers = 100
		const itemsPerWorker = 100

		// Concurrent Adds
		wg.Add(workers)
		for i := 0; i < workers; i++ {
			go func(start int) {
				defer wg.Done()
				for j := 0; j < itemsPerWorker; j++ {
					s.Add(start + j)
				}
			}(i * itemsPerWorker)
		}
		wg.Wait()

		if s.Len() != workers*itemsPerWorker {
			t.Errorf("Expected %d items, got %d", workers*itemsPerWorker, s.Len())
		}

		// Concurrent Contains
		wg.Add(workers)
		for i := 0; i < workers; i++ {
			go func(start int) {
				defer wg.Done()
				for j := 0; j < itemsPerWorker; j++ {
					if !s.Contains(start + j) {
						t.Errorf("Item %d not found", start+j)
					}
				}
			}(i * itemsPerWorker)
		}
		wg.Wait()

		// Concurrent Remove
		wg.Add(workers)
		for i := 0; i < workers; i++ {
			go func(start int) {
				defer wg.Done()
				for j := 0; j < itemsPerWorker; j++ {
					s.Remove(start + j)
				}
			}(i * itemsPerWorker)
		}
		wg.Wait()

		if s.Len() != 0 {
			t.Errorf("Expected empty set, got %d items", s.Len())
		}
	})

	t.Run("RangeConcurrency", func(t *testing.T) {
		s := NewSafeSet[int]()
		for i := 0; i < 100; i++ {
			s.Add(i)
		}

		var wg sync.WaitGroup
		const readers = 10
		wg.Add(readers)

		for i := 0; i < readers; i++ {
			go func() {
				defer wg.Done()
				s.Range(func(item int) bool {
					if !s.Contains(item) {
						t.Errorf("Item %d not found during range", item)
					}
					return true
				})
			}()
		}
		wg.Wait()
	})
}
