package wtype_test

import (
	"testing"

	"github.com/wuchieh/wtype"
)

func TestSet(t *testing.T) {
	t.Run("BasicOperations", func(t *testing.T) {
		s := wtype.NewSet[int]()
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

	t.Run("Duplicates", func(t *testing.T) {
		s := wtype.NewSet[string]()
		s.Add("a")
		s.Add("a")
		s.Add("a")

		if s.Len() != 1 {
			t.Error("Set should not allow duplicates")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		s := wtype.NewSet[float64]()
		s.Add(1.1)
		s.Add(2.2)
		s.Clear()

		if s.Len() != 0 {
			t.Error("Set should be empty after clear")
		}
	})

	t.Run("Range", func(t *testing.T) {
		s := wtype.NewSet[int]()
		s.Add(1)
		s.Add(2)
		s.Add(3)

		var sum int
		s.Range(func(i int) bool {
			sum += i
			return true
		})

		if sum != 6 {
			t.Error("Range should iterate all elements")
		}
	})
}
