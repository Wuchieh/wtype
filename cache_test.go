package wtype_test

import (
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

func TestNewCache(t *testing.T) {
	t.Run("create cache without initial data", func(t *testing.T) {
		c := wtype.NewCache[int](time.Second)
		if c == nil {
			t.Fatal("expected cache to be created")
		}
		if c.Get() != 0 {
			t.Errorf("expected default value 0, got %d", c.Get())
		}
	})

	t.Run("create cache with initial data", func(t *testing.T) {
		c := wtype.NewCache[int](time.Second, 42)
		if c.Get() != 42 {
			t.Errorf("expected 42, got %d", c.Get())
		}
	})

	t.Run("create cache with zero duration", func(t *testing.T) {
		c := wtype.NewCache[string](0, "test")
		if c.Get() != "test" {
			t.Errorf("expected 'test', got %s", c.Get())
		}
		time.Sleep(100 * time.Millisecond)
		if c.Get() != "test" {
			t.Error("cache should not expire with zero duration")
		}
	})
}

func TestCache_Set(t *testing.T) {
	t.Run("set and get value", func(t *testing.T) {
		c := wtype.NewCache[string](time.Second)
		c.Set("hello")
		if c.Get() != "hello" {
			t.Errorf("expected 'hello', got %s", c.Get())
		}
	})

	t.Run("set multiple times", func(t *testing.T) {
		c := wtype.NewCache[int](time.Second)
		c.Set(1)
		c.Set(2)
		c.Set(3)
		if c.Get() != 3 {
			t.Errorf("expected 3, got %d", c.Get())
		}
	})
}

func TestCache_Expiration(t *testing.T) {
	t.Run("data expires after duration", func(t *testing.T) {
		c := wtype.NewCache[int](50*time.Millisecond, 100)
		if c.Get() != 100 {
			t.Errorf("expected 100, got %d", c.Get())
		}

		time.Sleep(70 * time.Millisecond)
		if c.Get() != 0 {
			t.Errorf("expected data to be reset to 0, got %d", c.Get())
		}
	})

	t.Run("timer resets on Set", func(t *testing.T) {
		c := wtype.NewCache[string](50*time.Millisecond, "initial")

		time.Sleep(30 * time.Millisecond)
		c.Set("updated")

		time.Sleep(30 * time.Millisecond)
		if c.Get() != "updated" {
			t.Error("timer should have been reset, data should still be 'updated'")
		}

		time.Sleep(30 * time.Millisecond)
		if c.Get() != "" {
			t.Error("data should have expired after new timer duration")
		}
	})
}

func TestCache_StopTimer(t *testing.T) {
	t.Run("stop timer prevents expiration", func(t *testing.T) {
		c := wtype.NewCache[int](50*time.Millisecond, 42)
		c.StopTimer()

		time.Sleep(100 * time.Millisecond)
		if c.Get() != 42 {
			t.Errorf("data should not expire after stopping timer, got %d", c.Get())
		}
	})

	t.Run("stop timer on nil timer", func(t *testing.T) {
		c := wtype.NewCache[int](0)
		c.StopTimer() // should not panic
	})
}

func TestCache_ResetTimer(t *testing.T) {
	t.Run("reset timer extends expiration", func(t *testing.T) {
		c := wtype.NewCache[int](50*time.Millisecond, 99)

		time.Sleep(30 * time.Millisecond)
		c.ResetTimer()

		time.Sleep(30 * time.Millisecond)
		if c.Get() != 99 {
			t.Error("timer should have been reset, data should still exist")
		}

		time.Sleep(30 * time.Millisecond)
		if c.Get() != 0 {
			t.Error("data should have expired after reset timer duration")
		}
	})
}

func TestCache_SetDuration(t *testing.T) {
	t.Run("change duration", func(t *testing.T) {
		c := wtype.NewCache[int](100*time.Millisecond, 50)
		c.SetDuration(50 * time.Millisecond)
		c.Set(75) // This will use the new duration

		time.Sleep(70 * time.Millisecond)
		if c.Get() != 0 {
			t.Error("data should expire with new duration")
		}
	})

	t.Run("set duration to zero", func(t *testing.T) {
		c := wtype.NewCache[string](50*time.Millisecond, "test")
		c.SetDuration(0)
		c.Set("permanent")

		time.Sleep(100 * time.Millisecond)
		if c.Get() != "permanent" {
			t.Error("data should not expire with zero duration")
		}
	})
}

func TestCache_ComplexTypes(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	t.Run("cache with struct type", func(t *testing.T) {
		c := wtype.NewCache[User](time.Second)
		user := User{Name: "Alice", Age: 30}
		c.Set(user)

		got := c.Get()
		if got.Name != "Alice" || got.Age != 30 {
			t.Errorf("expected User{Alice, 30}, got %+v", got)
		}
	})

	t.Run("cache with pointer type", func(t *testing.T) {
		c := wtype.NewCache[*User](time.Second)
		user := &User{Name: "Bob", Age: 25}
		c.Set(user)

		got := c.Get()
		if got == nil || got.Name != "Bob" || got.Age != 25 {
			t.Errorf("expected &User{Bob, 25}, got %+v", got)
		}
	})

	t.Run("cache with slice type", func(t *testing.T) {
		c := wtype.NewCache[[]int](time.Second)
		data := []int{1, 2, 3, 4, 5}
		c.Set(data)

		got := c.Get()
		if len(got) != 5 || got[0] != 1 || got[4] != 5 {
			t.Errorf("expected [1,2,3,4,5], got %v", got)
		}
	})
}

func TestCache_ConcurrentAccess(t *testing.T) {
	t.Run("concurrent set and get", func(t *testing.T) {
		c := wtype.NewCache[int](time.Second, 0)

		done := make(chan bool)
		for i := 0; i < 10; i++ {
			go func(val int) {
				c.Set(val)
				_ = c.Get()
				done <- true
			}(i)
		}

		for i := 0; i < 10; i++ {
			<-done
		}
		// Test completes without panic/deadlock
	})
}

func TestCache_EdgeCases(t *testing.T) {
	t.Run("negative duration treated as immediate", func(t *testing.T) {
		c := wtype.NewCache[int](-time.Second, 42)
		if c.Get() == 0 {
			t.Error("cache should not expire with negative duration")
		}
	})

	t.Run("very long duration", func(t *testing.T) {
		c := wtype.NewCache[int](24*time.Hour, 100)
		if c.Get() != 100 {
			t.Error("cache should handle long durations")
		}
	})

	t.Run("multiple timer resets", func(t *testing.T) {
		c := wtype.NewCache[int](100*time.Millisecond, 1)
		for i := 0; i < 5; i++ {
			time.Sleep(20 * time.Millisecond)
			c.ResetTimer()
		}
		if c.Get() != 1 {
			t.Error("data should still exist after multiple resets")
		}
	})
}
