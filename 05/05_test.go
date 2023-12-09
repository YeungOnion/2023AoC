package main

import (
	"YeungOnion/2023AoC/avl"
	"YeungOnion/2023AoC/utils"
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
			value:    10,
			expected: 10,
		},
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
	}

	// construct sorted table
	seedTableInput := avl.NewBST[SeedMap](seedMapCompare)
	seedTableInput.Insert(SeedMap{src: 50, dest: 52, window: 48})
	seedTableInput.Insert(SeedMap{src: 98, dest: 50, window: 2})
	t.Log("var input seedTable:=\n", seedTableInput)

	for _, test := range tests {

		got := SeedTableEval(seedTableInput, test.value)
		if got != test.expected {
			t.Fatalf("input: %d\ngot %d, expected %d", test.value, got, test.expected)
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

func TestPartA(t *testing.T) {
	filename := "input.txt"
	expected := 226172555

	// stream file by words
	fs, closeHandle := utils.FileScanner(filename)
	defer closeHandle()
	fs.Split(bufio.ScanLines)

	seedNums := ScanSeedLine(fs)
	tables := ParseAllTables(fs)

	var minVal int
	for i, seed := range seedNums {
		val := PushThroughMaps(tables, seed)
		if i == 0 || val < minVal {
			minVal = val
		}
	}

	fmt.Println("min location: ", minVal)
	if got := minVal; expected != got {
		t.Fatalf("failed part a, expected %d, got %d", expected, got)
	}

	return
}
