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

func Scan[T any, R any, C ~[]T](collection C, iteratee func(R, T) R, base R) []R {
	result := make([]R, len(collection))
	prev := base
	for i := 0; i < len(collection); i++ {
		result[i] = iteratee(prev, collection[i])
		prev = result[i]
	}
	return result
}

func MapSeq[T any, R any, C ~[]T](collection C, iteratee func(T, int) R) []R {
	return lo.Map[T, R](collection, iteratee)
}

func Map[T any, R any, C ~[]T](collection C, iteratee func(T) R) []R {
	return lo.Map[T, R](collection, func(item T, _ int) R { return iteratee(item) })
}

func MapReduce[T any, R any, S any, C ~[]T](collection C, transform func(T) R, aggregator func(S, R) S, initial S) S {
	return lo.Reduce[R, S](
		Map[T, R, C](collection, transform),
		func(agg S, item R, _ int) S { return aggregator(agg, item) },
		initial,
	)
}

func MapReduceSeq[T any, R any, S any, C ~[]T](collection C, transform func(T, int) R, aggregator func(S, R, int) S, initial S) S {
	return lo.Reduce[R, S](lo.Map[T, R](collection, transform), aggregator, initial)
}
