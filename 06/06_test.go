package main

import (
	"testing"
)

func TestCounters(t *testing.T) {
	tests := []struct {
		name     string
		duration int
		record   int
		expected int
	}{
		{
			duration: 8,
			record:   8,
			expected: 5, // 2, 3, 4, 5, 6
		},
		{
			duration: 6,
			record:   8,
			expected: 1, // 1
		},
		{
			duration: 100,
			record:   2000,
			expected: 45,
		},
		{
			duration: 100,
			record:   200,
			expected: 95,
		},
		{
			duration: 80,
			record:   1200,
			expected: 39,
		},
		{
			name:     "sample partb",
			duration: 71530,
			record:   940200,
			expected: 71503,
		},
		{
			name:     "solution partb",
			duration: 62649190,
			record:   553101014731074,
			expected: 1,
		},
	}

	for _, test := range tests {
		if got := countSolns(test.duration, test.record); got != test.expected {
			t.Logf("Fail while testing,\n`%s': (%d, %d)", test.name, test.duration, test.record)
			t.Fatalf("got: %d, expected: %d", got, test.expected)
		}
	}
}
