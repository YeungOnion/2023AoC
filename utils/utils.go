package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// ScanWhile consumes input from scanner via Scan, if Scanned matches the predicate
// Then the text it appended to a string slice output
func ScanWhile(fs *bufio.Scanner, pred func(*bufio.Scanner) bool) []string {
	if pred == nil {
		pred = func(*bufio.Scanner) bool { return true }
	}

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

// ScanWhileString accumulates text until the predicatedScan is false
// and emits that text
func ScanWhileString(fs *bufio.Scanner, predicatedScan func(*bufio.Scanner) (string, bool)) []string {
	lines := make([]string, 0, 16)
	for {
		if s, ok := predicatedScan(fs); !ok {
			break
		} else {
			lines = append(lines, s)
		}
	}
	return lines
}

func numberSequenceBuffered(fs *bufio.Scanner) bool {
	re := regexp.MustCompile(`(\d+ *)+`)
	return re.MatchString(fs.Text())
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

// FileScanner opens a file and returns a scanner and a handle to close the file
func FileScanner(filename string) (*bufio.Scanner, func() error) {
	file, err := os.Open(filename)
	if err != nil {
		file.Close()
		panic(err)
	}
	return bufio.NewScanner(file), file.Close
}

// MustAtoi is used in tandem with lo.Map
func MustAtoi(s string, _ int) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}
