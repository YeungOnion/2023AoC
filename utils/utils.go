package utils

import (
	"bufio"
)

// scanWhile consumes input from scanner via Scan, if Scanned matches the predicate
// Then the text it appended to a string slice output
func ScanWhile(fs *bufio.Scanner, pred func(*bufio.Scanner) bool) []string {
	textRows := make([]string, 0, 16)
	for fs.Scan() {
		text := fs.Text()
		if !pred(fs) {
			break
		} else {
			textRows = append(textRows, text)
		}
	}
	return textRows
}

// All is a short circuit at first false of Reduce(iterable ~[]V, predicate func(V) bool) bool
func All[T ~[]V, V any](t T, pred func(V) bool) bool {
	for _, v := range t {
		if !pred(v) {
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
