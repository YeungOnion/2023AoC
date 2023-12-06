package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestMap(t *testing.T) {
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
