package main

import (
	"regexp"
	"testing"
)

func TestRegexpFwd(t *testing.T) {
	initBothRegex()

	tests := []struct {
		in  string
		exp []int
	}{
		{
			in:  "1abc2",
			exp: []int{0, 1},
		},
		{
			in:  "oneatwo3",
			exp: []int{0, 3},
		},
	}

	for _, test := range tests {
		got := regexp.MustCompile(forwardDigitRe).FindStringIndex(test.in)
		if len(got) == 0 {
			t.Fatalf("found no match in %s", test.in)
		} else if got[0] != test.exp[0] || got[1] != test.exp[1] {
			t.Fatalf("found match in position in %v, instead of %v", got, test.exp)
		}
	}
}

// func TestLineProcessDigitOnly(t *testing.T) {
// 	tests := []struct {
// 		in  string
// 		exp int
// 	}{
// 		{
// 			in:  "1abc2",
// 			exp: 12,
// 		},
// 		{
// 			in:  "oneatwo3",
// 			exp: 33,
// 		},
// 	}

// 	for _, test := range tests {
// 		got := extractNumberByDigitOrDigitName(test.in)
// 		if test.exp != got {
// 			t.Fatalf(
// 				"extractNumbers failed,\nin: %s\n got %d, expected %d",
// 				test.in,
// 				got,
// 				test.exp,
// 			)
// 		}
// 	}

// }

func TestLineProcessPartB(t *testing.T) {
	tests := []struct {
		in  string
		exp int
	}{
		{
			in:  "1abc2",
			exp: 12,
		},
		{
			in:  "oneatwo3",
			exp: 13,
		},
	}

	for _, test := range tests {
		got := extractNumberByDigitOrDigitName(test.in)
		if test.exp != got {
			t.Fatalf(
				"extractNumbers failed,\nin: %s\n got %d, expected %d",
				test.in,
				got,
				test.exp,
			)
		}
	}

}
