package main

import (
	"YeungOnion/2023AoC/set"
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		in  string
		exp set.Set[int]
	}{
		{
			in:  "12 1 244 23",
			exp: set.Set[int]{12: struct{}{}, 244: struct{}{}, 1: struct{}{}, 23: struct{}{}},
		},
		{
			in:  "1 2 | 5 6",
			exp: set.Set[int]{1: struct{}{}, 2: struct{}{}},
		},
	}

	for _, test := range tests {
		// read words from the string as buffer
		scan := bufio.NewScanner(strings.NewReader(test.in))
		scan.Split(bufio.ScanWords)

		// parser filtermap
		digitsRe := regexp.MustCompile(`\d+`)
		NumberString := func(s string) bool { return digitsRe.MatchString(s) }
		NumberStringer := func(s string) int {
			if val, err := strconv.Atoi(s); err == nil {
				return val
			} else {
				panic(err)
			}
		}

		// actual test
		got := ScanWhile(scan, NumberString, NumberStringer)
		if len(got) == 0 {
			t.Fatalf("consumed no elements in %s", test.in)
		} else if !got.Equiv(test.exp) {
			t.Fatalf("got %v, expected %v", got, test.exp)
		}
	}
}

