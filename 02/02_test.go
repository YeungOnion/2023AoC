package main

import (
	"testing"
)

func TestColorMax(t *testing.T) {
	tests := []struct {
		in    string
		color string
		exp   int
	}{
		{
			in:    "4 red 8 blue 8 red",
			color: "red",
			exp:   8,
		},
		{
			in:    "4 red 16 blue 5 blue 8 red; green",
			color: "blue",
			exp:   16,
		},
	}

	for _, test := range tests {
		got := findMaxOfColor(test.in, test.color)
		if got != test.exp {
			t.Fatalf("from %s\n got %d, expected %d", test.in, got, test.exp)
		}
	}

	return
}
