package iter

import (
	"github.com/samber/lo"
)

// All is a short circuit at first false of Reduce(iterable ~[]V, predicate func(V) bool) bool
func All[T ~[]V, V any](t T, pred func(V) bool) bool {
	for _, v := range t {
		if !pred(v) {
			return false
		}
	}
	return true
}

// All is a short circuit at first false of Reduce(iterable ~[]V, predicate func(V) bool) bool
func AllBool[T ~[]bool](t T) bool {
	for _, item := range t {
		if !item {
			return false
		}
	}
	return true
}

// Any is a short circuit at first true of Reduce(iterable ~[]V, predicate func(V) bool) bool
func Any[T ~[]V, V any](t T, pred func(V) bool) bool {
	for _, v := range t {
		b := pred(v)
		if b {
			return true
		}
	}
	return false
}

// Count returns the number of elements from the iterable that match the predicate
func Count[T ~[]V, V any](t T, pred func(V) bool) int {
	count := 0
	for _, v := range t {
		if pred(v) {
			count++
		}
	}
	return count
}

func Zip2With[T any, U any, R any](a []T, b []U, iteratee func(T, U) R) []R {
	return lo.Map[lo.Tuple2[T, U], R](
		lo.Zip2[T, U](a, b),
		func(tup lo.Tuple2[T, U], _ int) R {
			x, y := tup.Unpack()
			return iteratee(x, y)
		},
	)
}

func Scan[T any, R any](collection []T, iteratee func(R, T) R, base R) []R {
	result := make([]R, len(collection))
	prev := base
	for i := 0; i < len(collection); i++ {
		result[i] = iteratee(prev, collection[i])
		prev = result[i]
	}
	return result
}
