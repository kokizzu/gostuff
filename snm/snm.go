// Package snm provides convenience functions for slices and maps.
package snm

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Slice returns a new slice of size n whose values are the results
// of applying f on each index.
func Slice[T any](n int, f func(int) T) []T {
	s := make([]T, n)
	for i := range s {
		s[i] = f(i)
	}
	return s
}

// SliceToSlice returns a slice of the same length containing the results
// of applying f to the elements of s.
func SliceToSlice[A any, B any](a []A, f func(A) B) []B {
	b := make([]B, len(a))
	for i := range a {
		b[i] = f(a[i])
	}
	return b
}

// MapToMap returns a map containing the results of applying f to the key-value
// pairs of m.
// f should return a new key-value pair for the new map.
// Keys that appear more than once will override each other.
func MapToMap[K comparable, V any, K2 comparable, V2 any](
	m map[K]V, f func(K, V) (K2, V2)) map[K2]V2 {
	mm := make(map[K2]V2, len(m))
	for k, v := range m {
		k2, v2 := f(k, v)
		mm[k2] = v2
	}
	return mm
}

// FilterSlice returns a slice containing only the elements
// for which keep returns true.
func FilterSlice[S any](s []S, keep func(S) bool) []S {
	var result []S
	for _, e := range s {
		if keep(e) {
			result = append(result, e)
		}
	}
	return result
}

// FilterMap returns a map containing only the elements
// for which keep returns true.
func FilterMap[K comparable, V any](m map[K]V, keep func(K, V) bool) map[K]V {
	mm := map[K]V{}
	for k, v := range m {
		if keep(k, v) {
			mm[k] = v
		}
	}
	return mm
}

// Sorted sorts the input and returns it.
func Sorted[T constraints.Ordered](s []T) []T {
	slices.Sort(s)
	return s
}

// SortedFunc sorts the input and returns it.
func SortedFunc[T any](s []T, less func(T, T) bool) []T {
	slices.SortFunc(s, less)
	return s
}

// At returns the elements of t at the indexes in at.
func At[T any, I constraints.Integer](t []T, at []I) []T {
	result := make([]T, 0, len(at))
	for _, i := range at {
		result = append(result, t[i])
	}
	return result
}