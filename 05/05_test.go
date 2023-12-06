package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {

	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{
			value:    79,
			expected: 81,
		},
		{
			value:    98,
			expected: 50,
		},
		{
			value:    99,
			expected: 51,
		},
		{
			value:    100,
			expected: 100,
		},
		{
			value:    10,
			expected: 10,
		},
	}

	// construct sorted table
	input := seedTable{
		{src: 50, dest: 52, window: 48},
		{src: 98, dest: 50, window: 2},
	}

	for _, test := range tests {
		t.Logf("testing Eval(%d)\n", test.value)

		got := input.Eval(test.value)
		if got != test.expected {
			t.Logf("failed test %s,\nEval(%d)\n%v", test.name, test.value, input)
			t.Fatalf("got %d, expected %d", got, test.expected)
		}
	}
	return
}

func TestSampleA(t *testing.T) {
	fname := "sample-a.txt"
	expectedSeeds := "[79 14 55 13]"

	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	seedNums := ScanSeedLine(fileScanner)
	if fmt.Sprintf("%v", seedNums) != expectedSeeds {
		t.Fatalf("misread seed line as\ngot     : %v\nexpected: %v",
			seedNums, expectedSeeds)
	}

}
