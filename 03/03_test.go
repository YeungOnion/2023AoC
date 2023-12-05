package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAdjacentColumn(t *testing.T) {
	tests := []struct {
		in  string
		pos []int
		cmp string
		exp bool
	}{
		{
			cmp: ".......", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: false,
		},
		{
			cmp: "..123*.", // [][]int{{5,6}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: true,
		},
		{
			cmp: "*....?.", // [][]int{{0,1},{5,6}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: true,
		},
		{
			cmp: "......?", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: false,
		},
		{
			cmp: "?......", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: false,
		},
		{
			cmp: ".?.....", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: true,
		},
	}

	for _, test := range tests {
		// got := cal
		got := AdjacentColumn(test.pos, test.cmp, regexp.MustCompile(`[^\d\.]`))
		if got != test.exp {
			t.Fatalf("\n%s\n%s\n", test.in, test.cmp)
			t.Fatalf("got %v, expected %v\n", got, test.exp)
		}
	}
}

func TestGivenMiddleRow(t *testing.T) {
	tests := []struct {
		in       []string
		expected bool
	}{
		{
			in: []string{
				".......",
				"..123*.",
				"*....?.",
			},
			expected: true,
		},
		{
			in: []string{
				".......",
				"..123..",
				"*....?.",
			},
			expected: true,
		},
		{
			in: []string{
				".......",
				"..123..",
				"*.....?",
			},
			expected: false,
		},
		{
			in: []string{
				".......",
				"..123*.",
				".......",
			},
			expected: true,
		},
	}

	for _, test := range tests {
		// assert number of matches since testing
		index := GetTargetIndexes(test.in[1])[0]
		if HasNeighborPeripherals(index, test.in) != test.expected {
			out := fmt.Sprintf("expected match? = %v, assertion failed with input\n", test.expected)
			out += strings.Join(test.in, "\n")
			t.Fatal(out)
		}
	}

}
