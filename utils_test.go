package wtype_test

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/wuchieh/wtype"
)

func TestStringSlice(t *testing.T) {
	a := "abcdefg"

	if wtype.StringSlice(a, -1) != "g" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -2) != "fg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -3) != "efg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -4) != "defg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -5) != "cdefg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -6) != "bcdefg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -7) != "abcdefg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 0) != a {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 1) != "bcdefg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 2) != "cdefg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 3) != "defg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 4) != "efg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 5) != "fg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 6) != "g" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 1, 3) != "bc" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, 1, 99) != "bcdefg" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -2, -1) != "f" {
		t.Error("StringSlice error")
	}

	if wtype.StringSlice(a, -2, 99) != "fg" {
		t.Error("StringSlice error")
	}

	str := wtype.NewString("你好世界")

	if str.Slice(1, 2).String() != "好" {
		t.Error("StringSlice error")
	}
}

func TestStructStringTrim(t *testing.T) {
	type user struct {
		Name string
		Next *user
	}

	/*	toJson := func(u *user) string {
			b, _ := json.Marshal(u)
			return string(b)
		}
	*/
	u := &user{
		Name: " Alice ",
		Next: &user{
			Name: " Bob ",
			Next: &user{
				Name: " Charlie ",
				Next: &user{
					Name: " Dave ",
				},
			},
		},
	}

	wtype.StructStringTrim(&u)

	tempU := u

	for tempU != nil {
		if strings.Contains(tempU.Name, " ") {
			t.Error("StructStringTrim error")
		}

		tempU = tempU.Next
	}

	u2 := &user{
		Name: " Alice ",
	}

	u2.Next = u2

	wtype.StructStringTrim(&u2)

	if u2 != u2.Next || u2.Next.Next.Name != "Alice" {
		t.Error("StructStringTrim error")
	}
}

func TestSliceToMap(t *testing.T) {
	type user struct {
		ID   int
		Name string
		Age  int
	}

	users := []user{
		{
			ID:   1,
			Name: "Alice",
			Age:  35,
		},
		{
			ID:   2,
			Name: "Bob",
			Age:  28,
		},
		{
			ID:   3,
			Name: "Charlie",
			Age:  14,
		},
	}

	m := wtype.SliceToMap(users, func(i int, u user) int {
		return u.ID
	})

	if m[1].Name != "Alice" {
		t.Error("SliceToMap error")
	}

	if m[2].Name != "Bob" {
		t.Error("SliceToMap error")
	}

	if m[3].Name != "Charlie" {
		t.Error("SliceToMap error")
	}
}

func TestSlicePointConvert(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	slice2 := wtype.SliceUnPointConvert(wtype.SlicePointConvert(slice))

	for i := 0; i < len(slice2); i++ {
		if slice2[i] != slice[i] {
			t.Error("SlicePointConvert error")
		}
	}
}

func TestSliceUnPointConvert(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	slice2 := wtype.SliceUnPointConvert(wtype.SlicePointConvert(slice))

	for i := 0; i < len(slice2); i++ {
		if slice2[i] != slice[i] {
			t.Error("SlicePointConvert error")
		}
	}
}

func TestDoShared(t *testing.T) {
	type data struct {
		do    int
		start int
		end   int
	}
	var result data
	key := time.Now().String()
	fn := func() int {
		num, err := wtype.DoShared(key, func() (int, error) {
			result.do++
			time.Sleep(time.Second)
			return 0, nil
		})
		if err != nil {
			return 0
		}
		return num
	}

	var wg sync.WaitGroup
	runTime := 10
	wg.Add(runTime)
	for i := 0; i < runTime; i++ {
		go func() {
			defer wg.Done()
			result.start++
			fn()
			result.end++
		}()
	}

	wg.Wait()

	if result.do != 1 {
		t.Error("do time error", result.do)
	}

	if result.start != runTime {
		t.Error("start time error")
	}

	if result.end != runTime {
		t.Error("end time error")
	}
}

func TestDoSharedChan(t *testing.T) {
	key := time.Now().String()

	fn := func() <-chan wtype.SharedChanResult[int] {
		return wtype.DoSharedChan(key, func() (int, error) {
			time.Sleep(time.Second)
			return 0, nil
		})
	}

	var wg sync.WaitGroup
	runTime := 10
	wg.Add(runTime)
	for i := 0; i < runTime; i++ {
		go func() {
			defer wg.Done()
			result := <-fn()
			if (result.Val + 1) != 1 {
				t.Error("DoSharedChan error")
			}
		}()
	}

	wg.Done()
}

func TestDoSharedForget(t *testing.T) {
	type data struct {
		do    int
		start int
		end   int
	}
	var result data
	key := time.Now().String()
	fn := func() int {
		num, err := wtype.DoShared(key, func() (int, error) {
			result.do++
			time.Sleep(time.Second)
			return 0, nil
		})
		if err != nil {
			return 0
		}
		return num
	}

	var wg sync.WaitGroup
	runTime := 10
	wg.Add(runTime)
	for i := 0; i < runTime; i++ {
		go func() {
			defer wg.Done()
			result.start++
			wtype.DoSharedForget(key)
			fn()
			result.end++
		}()

		time.Sleep(time.Millisecond)
	}

	wg.Wait()

	if result.do <= 1 {
		t.Error("do time error", result.do)
	}

	if result.start != runTime {
		t.Error("start time error")
	}

	if result.end != runTime {
		t.Error("end time error")
	}
}
