package snm

import (
	"fmt"
	"slices"
	"testing"

	"golang.org/x/exp/maps"
)

func TestSlice(t *testing.T) {
	want := []int{1, 4, 9, 16}
	got := Slice(4, func(i int) int { return (i + 1) * (i + 1) })
	if !slices.Equal(got, want) {
		t.Fatalf("Slice((i+1)*(i+1))=%v, want %v", got, want)
	}
}

func TestSliceToSlice(t *testing.T) {
	input := []int{1, 4, 9, 16}
	want := []float64{1.5, 4.5, 9.5, 16.5}
	got := SliceToSlice(input, func(i int) float64 {
		return float64(i) + 0.5
	})
	if !slices.Equal(got, want) {
		t.Fatalf("SliceToSlice(%v)=%v, want %v", input, got, want)
	}
}

func TestMapToMap(t *testing.T) {
	input := map[string]string{"a": "bbb", "cccc": "ddddddd"}
	want := map[int]int{1: 3, 4: 7}
	got := MapToMap(input, func(k, v string) (int, int) {
		return len(k), len(v)
	})
	if !maps.Equal(got, want) {
		t.Fatalf("MapToMap(%v)=%v, want %v", input, got, want)
	}
}

func TestMapToMap_equalKeys(t *testing.T) {
	input := map[string]string{"a": "bbb", "cccc": "ddddddd", "e": "ff"}
	want1 := map[int]int{1: 3, 4: 7}
	want2 := map[int]int{1: 2, 4: 7}
	got := MapToMap(input, func(k, v string) (int, int) {
		return len(k), len(v)
	})
	if !maps.Equal(got, want1) && !maps.Equal(got, want2) {
		t.Fatalf("MapToMap(%v)=%v, want %v or %v", input, got, want1, want2)
	}
}

func TestDefaultMap(t *testing.T) {
	m := NewDefaultMap[int, string](func(i int) string {
		return fmt.Sprint(i + 1)
	})
	if got, want := m.Get(2), "3"; got != want {
		t.Fatalf("Get(%d)=%s, want %s", 2, got, want)
	}
	if got, want := m.Get(6), "7"; got != want {
		t.Fatalf("Get(%d)=%s, want %s", 6, got, want)
	}
	m.Set(2, "a")
	if got, want := m.Get(2), "a"; got != want {
		t.Fatalf("Get(%d)=%s, want %s", 2, got, want)
	}
	if got, want := m.Get(6), "7"; got != want {
		t.Fatalf("Get(%d)=%s, want %s", 6, got, want)
	}
	if got, want := len(m.M), 2; got != want {
		t.Fatalf("Len=%d, want %d", got, want)
	}
}

func TestCompare(t *testing.T) {
	input := []int{3, 4, 2, 1, 5}
	want := []int{1, 2, 3, 4, 5}
	wantr := []int{5, 4, 3, 2, 1}

	cp := slices.Clone(input)
	slices.SortFunc(cp, Compare)
	if !slices.Equal(cp, want) {
		t.Errorf("SortFunc(%v, Compare)=%v, want %v",
			input, cp, want)
	}

	cp = slices.Clone(input)
	slices.SortFunc(cp, CompareReverse)
	if !slices.Equal(cp, wantr) {
		t.Errorf("SortFunc(%v, Compare)=%v, want %v",
			input, cp, wantr)
	}
}

func ExampleSortedKeys() {
	ages := map[string]int{
		"Alice":   30,
		"Bob":     20,
		"Charlie": 25,
	}
	for _, name := range SortedKeys(ages) {
		fmt.Printf("%s: %d\n", name, ages[name])
	}
	// Output:
	// Bob: 20
	// Charlie: 25
	// Alice: 30
}

func ExampleSortedKeysFunc_reverse() {
	ages := map[string]int{
		"Alice":   30,
		"Bob":     20,
		"Charlie": 25,
	}
	// Sort by reverse natural order.
	for _, name := range SortedKeysFunc(ages, CompareReverse) {
		fmt.Printf("%s: %d\n", name, ages[name])
	}
	// Output:
	// Alice: 30
	// Charlie: 25
	// Bob: 20
}
