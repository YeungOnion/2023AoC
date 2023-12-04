package main

import "testing"

func TestAdjacentColumn(t *testing.T) {
	tests := []struct {
		in  string
		cmp string
		pos []int
		exp bool
	}{
		{
			in:  "..123*.",
			cmp: "..123*.",
			pos: []int{2, 5},
			exp: true,
		},
		{
			in:  "..123*.",
			cmp: "*....?.",
			pos: []int{2, 5},
			exp: true,
		},
		{
			in:  "..123*.",
			cmp: "*.....?", // [][]int {{0,1},{6,7}}
			pos: []int{2, 5},
			exp: false,
		},
	}

	for _, test := range tests {
		// got := cal
		got := AdjacentColumn(test.pos, test.cmp)
		if got != test.exp {
			t.Fatalf("\n%s\n%s\n", test.in, test.cmp)
			t.Fatalf("got %v, expected %v\n", got, test.exp)
		}
	}

}
