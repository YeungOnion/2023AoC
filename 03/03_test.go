package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func TestHasNeighborsPartA(t *testing.T) {
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

	periphRe := regexp.MustCompile(`[^\d\.]`)
	targetRe := regexp.MustCompile(`\d+`)
	for _, test := range tests {
		// assert number of matches since testing
		index := GetTargetIndexes(test.in[1], targetRe)[0]
		if HasNeighborPeripherals(index, test.in, periphRe) != test.expected {
			out := fmt.Sprintf("expected match? = %v, assertion failed with input\n", test.expected)
			out += strings.Join(test.in, "\n")
			t.Fatal(out)
		}
	}

}

func TestCountNeighborsPartB(t *testing.T) {
	tests := []struct {
		in       []string
		expected int
	}{
		{
			in: []string{
				".......",
				"..123*.",
				".....?.",
			},
			expected: 1,
		},
		{
			in: []string{
				".......",
				"*....?.",
				"..123..",
			},
			expected: 0,
		},
		{
			in: []string{
				".132...",
				"*.....?",
				"..123..",
			},
			expected: 1,
		},
		{
			in: []string{
				"..123..",
				"..123*.",
				".......",
			},
			expected: 2,
		},
		{
			in: []string{
				"..123..",
				"..123*.",
				"..123..",
			},
			expected: 3,
		},
	}

	periphRe := regexp.MustCompile(`\d+`)
	targetRe := regexp.MustCompile(`\*`)
	for _, test := range tests {
		// assert number of matches since testing
		index := GetTargetIndexes(test.in[1], targetRe)[0]
		got := CountNeighborPeripherals(index, test.in, periphRe)
		if got != test.expected {
			out := fmt.Sprintf("expected %d, got %d\n", test.expected, got)
			out += strings.Join(test.in, "\n")
			t.Fatal(out)
		} else {
			t.Log("passed case")
		}
	}

}

func TestSamplePartA(t *testing.T) {
	exp := 4361
	infile := "sample.txt"

	// stream file by words
	file, err := os.Open(infile)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	gotPart, _ := ScanAndScore(fileScanner)
	if exp != gotPart {
		t.Fatalf("expected %d, got %d", exp, gotPart)
	}

}

func TestPartA(t *testing.T) {
	exp := 539590
	infile := "input.txt"

	// stream file by words
	file, err := os.Open(infile)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	gotPart, _ := ScanAndScore(fileScanner)
	if exp != gotPart {
		t.Fatalf("expected %d, got %d", exp, gotPart)
	}
}
